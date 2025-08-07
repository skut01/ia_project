<script setup>
defineProps({
  content: String,
  role: {
    type: String,
    default: 'user'
  }
})
</script>

<template>
  <div
    class="flex "
    :class="role === 'assistant' ? 'justify-start' : 'justify-end'"
  >
    <div
      :class="[
        'p-3 max-w-[75%] rounded-xl shadow-sm drop-shadow-sm',
        role === 'assistant' ? 'bg-sky-100 text-left' : 'bg-emerald-100 text-right'
      ]"
    >
      <p class="text-sm text-gray-700 whitespace-pre-wrap">{{ content }}</p>
      <span class="text-xs text-gray-500 block mt-1">
        {{ role === 'assistant' ? 'Assistant' : 'Utilisateur' }}
      </span>
    </div>
  </div>
</template>
