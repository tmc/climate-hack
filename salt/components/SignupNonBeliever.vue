<!--
  This example requires some changes to your config:
  
  ```
  // tailwind.config.js
  module.exports = {
    // ...
    plugins: [
      // ...
      require('@tailwindcss/forms'),
    ],
  }
  ```
-->
<template>
  <form @submit.prevent="submitForm">
    <div class="space-y-12">
      <div class="border-b border-white/10 pb-12">
        <h2 class="text-base font-semibold leading-7 text-white">
          Tell Your Friends about Salt Reactors
        </h2>
        <p class="mt-1 text-sm leading-6 text-gray-400">
          Let's get you started educating the masses.
        </p>

        <div class="mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6">
          <div class="sm:col-span-4">
            <label
              for="name"
              class="block text-sm font-medium leading-6 text-white"
              ><span class="line-through">Victim</span> Open-Minded
              Non-Believer</label
            >
            <div class="mt-2">
              <div
                class="flex rounded-md bg-white/5 ring-1 ring-inset ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-red-500"
              >
                <input
                  v-model="formData.name"
                  type="text"
                  name="name"
                  id="name"
                  autocomplete="name"
                  class="flex-1 border-0 bg-transparent py-1.5 pl-1 text-white focus:ring-0 sm:text-sm sm:leading-6"
                  placeholder="Mr. Thorium"
                />
              </div>
            </div>
          </div>

          <div class="sm:col-span-4">
            <label
              for="phone"
              class="block text-sm font-medium leading-6 text-white"
              >Phone Number of the Charlatan</label
            >
            <div class="mt-2">
              <div
                class="flex rounded-md bg-white/5 ring-1 ring-inset ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-red-500"
              >
                <input
                  v-model="formData.phone"
                  type="number"
                  name="phone"
                  id="phone"
                  autocomplete="phone"
                  class="flex-1 border-0 bg-transparent py-1.5 pl-1 text-white focus:ring-0 sm:text-sm sm:leading-6"
                  placeholder="+1234567890"
                />
              </div>
            </div>
          </div>

          <div class="col-span-full">
            <label
              for="other_info"
              class="block text-sm font-medium leading-6 text-white"
              >Special Considerations / Information</label
            >
            <div class="mt-2">
              <textarea
                v-model="formData.other_info"
                id="other_info"
                name="other_info"
                rows="3"
                class="block w-full rounded-md border-0 bg-white/5 py-1.5 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-red-500 sm:text-sm sm:leading-6"
              />
            </div>
            <p class="mt-3 text-sm leading-6 text-gray-400">
              Write a few sentences about your target to really personalize this
              appeal to reason.
            </p>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-6 flex items-center justify-end gap-x-6">
      <button
        type="button"
        @click="notAllowed()"
        class="text-sm font-semibold leading-6 text-white"
      >
        Cancel
      </button>
      <button
        type="submit"
        class="rounded-md bg-red-500 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-red-500"
      >
        Start Conversation
      </button>
    </div>
  </form>
</template>

<script setup>
import { PhotoIcon, UserCircleIcon } from "@heroicons/vue/24/solid";
import { ref } from "vue";

// function that opens popup that says "not allowed"
const notAllowed = () => {
  alert("The truth cannot be stifled by the likes of you.");
};

const formData = ref({
  phone: "",
  name: "",
  other_info: "",
});

const emit = defineEmits(["form-submit"]);

const submitForm = () => {
  // Do validation or other processing here
  emit("form-submit", formData.value);
};
</script>
