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
  <form>
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
              for="username"
              class="block text-sm font-medium leading-6 text-white"
              ><span class="line-through">Victim</span> Open-Minded
              Non-Believer</label
            >
            <div class="mt-2">
              <div
                class="flex rounded-md bg-white/5 ring-1 ring-inset ring-white/10 focus-within:ring-2 focus-within:ring-inset focus-within:ring-red-500"
              >
                <input
                  type="text"
                  name="username"
                  id="username"
                  autocomplete="username"
                  class="flex-1 border-0 bg-transparent py-1.5 pl-1 text-white focus:ring-0 sm:text-sm sm:leading-6"
                  placeholder="+1234567890"
                />
              </div>
            </div>
          </div>

          <div class="col-span-full">
            <label
              for="about"
              class="block text-sm font-medium leading-6 text-white"
              >Special Considerations / Information</label
            >
            <div class="mt-2">
              <textarea
                id="about"
                name="about"
                rows="3"
                class="block w-full rounded-md border-0 bg-white/5 py-1.5 text-white shadow-sm ring-1 ring-inset ring-white/10 focus:ring-2 focus:ring-inset focus:ring-red-500 sm:text-sm sm:leading-6"
              />
            </div>
            <p class="mt-3 text-sm leading-6 text-gray-400">
              Write a few sentences about your target to really personalize this
              appeal to reason.
            </p>
          </div>

          <!-- <div class="col-span-full">
            <label
              for="cover-photo"
              class="block text-sm font-medium leading-6 text-white"
              >Photo of the Non-Believer</label
            >
            <div
              class="mt-2 flex justify-center rounded-lg border border-dashed border-white/25 px-6 py-10"
            >
              <div class="text-center">
                <PhotoIcon
                  class="mx-auto h-12 w-12 text-gray-500"
                  aria-hidden="true"
                />
                <div class="mt-4 flex text-sm leading-6 text-gray-400">
                  <label
                    for="file-upload"
                    class="relative cursor-pointer rounded-md bg-gray-900 font-semibold text-white focus-within:outline-none focus-within:ring-2 focus-within:ring-red-600 focus-within:ring-offset-2 focus-within:ring-offset-gray-900 hover:text-red-500"
                  >
                    <span>Upload a file</span>
                    <input
                      id="file-upload"
                      name="file-upload"
                      type="file"
                      class="sr-only"
                    />
                  </label>
                  <p class="pl-1">or drag and drop</p>
                </div>
                <p class="text-xs leading-5 text-gray-400">
                  PNG, JPG, GIF up to 10MB
                </p>
              </div>
            </div>
          </div> -->
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

// function that opens popup that says "not allowed"
const notAllowed = () => {
  alert(
    "The truth cannot be stifled by the likes of you. Go through with it or perish a coward."
  );
};
</script>
