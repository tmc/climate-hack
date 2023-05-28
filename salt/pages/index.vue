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
                      <h1
                        @click="active_page == 'welcome'"
                        class="text-xl text-white font-bold"
                      >
                        SALT
                      </h1>
                    </div>
                    <nav class="flex flex-1 flex-col">
                      <ul role="list" class="flex flex-1 flex-col gap-y-7">
                        <li>
                          <ul role="list" class="-mx-2 space-y-1">
                            <li>
                              <button @click="active_page = 'about'">
                                About Salt Reactors
                              </button>
                            </li>
                            <li>
                              <button @click="active_page = 'form'">
                                Start Convincing
                              </button>
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
          <div
            class="flex grow flex-col gap-y-5 overflow-y-auto bg-black/10 px-6 ring-1 ring-white/5"
          >
            <div class="flex h-16 shrink-0 items-center">
              <button
                @click="active_page == 'welcome'"
                class="text-xl text-white font-bold"
              >
                SALT
              </button>
            </div>
            <nav class="flex flex-1 flex-col">
              <ul role="list" class="flex flex-1 flex-col gap-y-7">
                <li>
                  <ul role="list" class="-mx-2 space-y-1">
                    <li>
                      <button @click="active_page = 'about'">
                        About Salt Reactors
                      </button>
                    </li>
                    <li>
                      <button @click="active_page = 'form'">
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
          <div v-if="active_page == 'welcome'" class="">
            <Welcome />
          </div>

          <!-- About -->
          <div
            v-else-if="active_page == 'about'"
            class="max-w-4xl mx-auto pt-12"
          >
            <AboutSaltReactors />
          </div>

          <!-- Form -->
          <div
            v-else-if="active_page == 'form'"
            class="max-w-4xl mx-auto pt-12"
          >
            <SignupNonBeliever @form-submit="handleFormSubmit" />
          </div>

          <!-- Conversation -->
          <div
            v-else-if="active_page == 'conversation'"
            class="max-w-4xl mx-auto pt-12"
          >
            <Conversation />
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

// const query = gql`
//   query {
//     me {
//       isActive
//       grants {
//         title
//         description
//         id
//         ownerId
//       }
//     }
//   }
// `;
// const variables = { limit: 5 };

// const { data } = await useAsyncQuery(query, variables);

const sendTextMutation = gql`
  mutation informNonBeliver($phone: String!) {
    informNonBeliver(input: { phone: $phone }) {
      conversation {
        id
      }
    }
  }
`;

const { mutate: sendText } = useMutation(sendTextMutation);

// form_active
const active_page = ref("welcome");

// Other data here
let submittedFormData = ref(null);

const handleFormSubmit = async (formData) => {
  submittedFormData.value = formData;
  const variables = {
    phone: formData.phone,
  };
  sendText(variables);
  active_page.value = "conversation";
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
