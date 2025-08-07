<script setup>
import { ref, onMounted } from 'vue'
import { fetchConversation } from '@/services/conversationGetterService'
import MessageBubble from './MessageBubble.vue'

// Props
const props = defineProps({
  convId: {
    type: String,
    required: true
  }
})

// Data
const messages = ref([])
const error = ref(null)
const loading = ref(true)

// Fetch data
onMounted(async () => {
  try {
    const res = await fetchConversation(props.convId)
    console.log('messages reçus :', res)
    messages.value = res
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div class="flex flex-col gap-4 px-6 py-4 w-full max-w-screen-xl mx-auto">
    <div v-if="loading" class="text-gray-500">Chargement...</div>
    <div v-else-if="error" class="text-red-500">Erreur : {{ error }}</div>
    <div v-else class="flex flex-col gap-3">
      <MessageBubble
        v-for="(msg, index) in messages"
        :key="index"
        :content="msg.Content"
        :role="msg.Role"
      />
    </div>
  </div>
</template>
