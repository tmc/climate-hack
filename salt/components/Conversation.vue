<template>
  <div class="px-8">
    <h2 class="text-base font-semibold leading-7 text-white">
      ThoriumGPT's Conversation with x
    </h2>
    <!-- <p class="mt-1 text-sm leading-6 text-gray-400">
          lorem
        </p> -->
    <div class="flex flex-col w-full justify-center space-y-4 mt-8">
      <div class="grid grid-cols-6">
        <div class="bg-gray-700 px-4 py-2 col-span-4 rounded">
          <h4 class="text-sm text-gray-400">user</h4>
          <p>Lorem Ipsum Dolor sit salt reactor.</p>
        </div>
      </div>
      <div class="grid grid-cols-6">
        <div class="col-start-2 col-span-4 bg-gray-800 px-4 py-2 rounded">
          <h4 class="text-sm text-gray-400">ThoriumGPT</h4>
          <p>Lorem Ipsum Dolor sit salt reactor.</p>
          ID: {{ id }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { defineProps } from 'vue'

  const props = defineProps({
  id: {
    type: String,
    required: true,
  },
})
</script>