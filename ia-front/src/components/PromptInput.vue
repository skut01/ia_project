<script setup>
import { ref } from 'vue'

const props = defineProps({
  loading: Boolean
})

const emit = defineEmits(['submit'])

const input = ref('')

function handleSubmit() {
    if (props.loading || input.value.trim() === '') return
    if (input.value.trim() !== '') {
        emit('submit', input.value.trim())
        input.value = ''
    }
}
</script>


<template>
  <form @submit.prevent="handleSubmit" class="flex items-center gap-3 p-4 bg-neutral-800 rounded-t-lg">
    <!-- loader -->
        <div class="w-5 h-5 flex items-center justify-center">
      <svg
        v-if="loading"
        class="animate-spin text-emerald-400 w-5 h-5"
        fill="none"
        viewBox="0 0 24 24"
      >
        <circle
          class="opacity-25"
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          stroke-width="4"
        />
        <path
          class="opacity-75"
          fill="currentColor"
          d="M4 12a8 8 0 018-8v8H4z"
        />
      </svg>

      <svg
        v-else
        class="text-emerald-400 w-5 h-5"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        stroke-width="2"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M5 13l4 4L19 7"
        />
      </svg>
    </div>

    <!-- input -->
    <input
      v-model="input"
      type="text"
      placeholder="Écrivez votre message..."
      class="flex-1 px-4 py-2 rounded-md bg-neutral-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-emerald-500"
    />

    <!-- bouton -->
    <button
        type="submit"
        :disabled="loading"
        :class="[
            'px-4 py-2 rounded-md transition',
            loading
            ? 'bg-gray-600 cursor-not-allowed text-gray-300'
            : 'bg-emerald-600 text-white hover:bg-emerald-700'
        ]"
        >
        Envoyer
    </button>

  </form>
</template>

