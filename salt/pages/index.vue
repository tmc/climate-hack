<template>
  <div class="text-white min-h-screen">
    <div>
      <div>
        <TransitionRoot as="template" :show="sidebarOpen">
          <dialog
            as="div"
            class="relative z-50 xl:hidden"
            @close="sidebarOpen = false"
          >
            <TransitionChild
              as="template"
              enter="transition-opacity ease-linear duration-300"
              enter-from="opacity-0"
              enter-to="opacity-100"
              leave="transition-opacity ease-linear duration-300"
              leave-from="opacity-100"
              leave-to="opacity-0"
            >
              <div class="fixed inset-0 bg-gray-900/80" />
            </TransitionChild>

            <div class="fixed inset-0 flex">
              <TransitionChild
                as="template"
                enter="transition ease-in-out duration-300 transform"
                enter-from="-translate-x-full"
                enter-to="translate-x-0"
                leave="transition ease-in-out duration-300 transform"
                leave-from="translate-x-0"
                leave-to="-translate-x-full"
              >
                <DialogPanel class="relative mr-16 flex w-full max-w-xs flex-1">
                  <TransitionChild
                    as="template"
                    enter="ease-in-out duration-300"
                    enter-from="opacity-0"
                    enter-to="opacity-100"
                    leave="ease-in-out duration-300"
                    leave-from="opacity-100"
                    leave-to="opacity-0"
                  >
                    <div
                      class="absolute left-full top-0 flex w-16 justify-center pt-5"
                    >
                      <button
                        type="button"
                        class="-m-2.5 p-2.5"
                        @click="sidebarOpen = false"
                      >
                        <span class="sr-only">Close sidebar</span>
                        <XMarkIcon
                          class="h-6 w-6 text-white"
                          aria-hidden="true"
                        />
                      </button>
                    </div>
                  </TransitionChild>
                  <!-- Sidebar component, swap this element with another sidebar if you like -->
                  <div
                    class="flex grow flex-col gap-y-5 overflow-y-auto bg-gray-900 px-6 ring-1 ring-white/10"
                  >
                    <div class="flex h-16 shrink-0 items-center">
                      <!-- <img
                        class="h-8 w-auto"
                        src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
                        alt="Your Company"
                      /> -->
                      <h1 class="text-xl text-white font-bold">
                        Grant Accountability
                      </h1>
                    </div>
                    <nav class="flex flex-1 flex-col">
                      <ul role="list" class="flex flex-1 flex-col gap-y-7">
                        <li>
                          <ul role="list" class="-mx-2 space-y-1">
                            <li v-for="item in navigation" :key="item.name">
                              <a
                                :href="item.href"
                                :class="[
                                  item.current
                                    ? 'bg-gray-800 text-white'
                                    : 'text-gray-400 hover:text-white hover:bg-gray-800',
                                  'group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold',
                                ]"
                              >
                                <component
                                  :is="item.icon"
                                  class="h-6 w-6 shrink-0"
                                  aria-hidden="true"
                                />
                                {{ item.name }}
                              </a>
                            </li>
                          </ul>
                        </li>
                      </ul>
                    </nav>
                  </div>
                </DialogPanel>
              </TransitionChild>
            </div>
          </dialog>
        </TransitionRoot>

        <!-- Static sidebar for desktop -->
        <div
          class="hidden xl:fixed xl:inset-y-0 xl:z-50 xl:flex xl:w-72 xl:flex-col"
        >
          <!-- Sidebar component, swap this element with another sidebar if you like -->
          <div
            class="flex grow flex-col gap-y-5 overflow-y-auto bg-black/10 px-6 ring-1 ring-white/5"
          >
            <div class="flex h-16 shrink-0 items-center">
              <!-- <img
                class="h-8 w-auto"
                src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=500"
                alt="Your Company"
              /> -->
              <button
                @click="
                  about_active = false;
                  form_active = false;
                  welcome_active = true;
                "
                class="text-xl text-white font-bold"
              >
                SALT
              </button>
            </div>
            <nav class="flex flex-1 flex-col">
              <ul role="list" class="flex flex-1 flex-col gap-y-7">
                <li>
                  <ul role="list" class="-mx-2 space-y-1">
                    <li v-for="item in navigation" :key="item.name">
                      <a
                        :href="item.href"
                        :class="[
                          item.current
                            ? 'bg-gray-800 text-white'
                            : 'text-gray-400 hover:text-white hover:bg-gray-800',
                          'group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold',
                        ]"
                      >
                        <component
                          :is="item.icon"
                          class="h-6 w-6 shrink-0"
                          aria-hidden="true"
                        />
                        {{ item.name }}
                      </a>
                    </li>
                    <li>
                      <button
                        @click="
                          about_active = true;
                          form_active = false;
                          welcome_active = false;
                        "
                      >
                        About Salt Reactors
                      </button>
                    </li>
                    <li>
                      <button
                        @click="
                          form_active = true;
                          about_active = false;
                          welcome_active = false;
                        "
                      >
                        Start Convincing
                      </button>
                    </li>
                  </ul>
                </li>
              </ul>
            </nav>
          </div>
        </div>

        <transition name="slide-fade" mode="out-in">
          <!-- Welcome -->
          <div v-if="welcome_active" class="">
            <Welcome />
          </div>

          <!-- About -->
          <div v-else-if="about_active" class="max-w-4xl mx-auto pt-12">
            <AboutSaltReactors />
          </div>

          <!-- Form -->
          <div v-else-if="form_active" class="max-w-4xl mx-auto pt-12">
            <SignupNonBeliever @form-submit="handleFormSubmit" />
          </div>
        </transition>

        <!-- <Transition>
      <div>
        <div class="sliding-background mt-64"></div>
      </div>
    </Transition> -->
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const query = gql`
  query {
    me {
      isActive
      grants {
        title
        description
        id
        ownerId
      }
    }
  }
`;
const variables = { limit: 5 };

const { data } = await useAsyncQuery(query, variables);

// form_active
const welcome_active = ref(true);
const about_active = ref(false);
const form_active = ref(false);

// Other data here
let submittedFormData = ref(null);

const handleFormSubmit = (formData) => {
  submittedFormData.value = formData;
  // Do whatever you need with the data here
  console.log(formData);
};
</script>

<style scoped>
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.5s ease;
}

.slide-fade-enter-from {
  opacity: 0;
  transform: translateX(-10%);
}

.slide-fade-leave-to {
  opacity: 0;
  transform: translateX(10%);
}

.slide-fade-enter-to,
.slide-fade-leave-from {
  opacity: 1;
  transform: translateX(0);
}

.sliding-background {
  background: url("/salt.png") repeat-x;
  height: 560px;
  width: 7076px;
  animation: slide 30s linear infinite;
}

@keyframes slide {
  0% {
    transform: translate3d(0, 0, 0);
  }
  100% {
    transform: translate3d(-1692px, 0, 0);
  }
}
</style>
