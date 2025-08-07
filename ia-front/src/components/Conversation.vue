<script setup>
import { fetchConversation } from '../services/conversationGetterService'
import { sendMessageToChat } from '../services/messageSender'
import MessageBubble from './MessageBubble.vue'
import PromptInput from './PromptInput.vue'
import { ref, onMounted, watch,nextTick  } from 'vue'

const props = defineProps({
  convId: {
    type: String,
    required: true,
  },
})

const isThinking = ref(false)
const messages = ref([])
const loading = ref(true)
const error = ref(null)
const model = 'openai/gpt-oss-20b:free'

function scrollToBottom() {
  nextTick(() => {
    if (chatContainer.value) {
      chatContainer.value.scrollTop = chatContainer.value.scrollHeight
    }
  })
}

onMounted(async () => {
  try {
    messages.value = await fetchConversation(props.convId)
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
  scrollToBottom()
})

const chatContainer = ref(null)



async function handlePrompt(promptText) {
  const userMessage = {
    Role: 'user',
    Content: promptText,
  }

  // Ajoute le message de l'utilisateur
  messages.value.push(userMessage)
  scrollToBottom()
  isThinking.value = true
  try {
    const responseFromModel = await sendMessageToChat({
      conversation_id: Number(props.convId),
      model,
      message: userMessage,
    })

    console.log('Réponse reçue :', responseFromModel)
    const assistantMessage = {
      Role: 'assistant',
      Content: responseFromModel.response || 'Réponse vide',
    }

    messages.value.push(assistantMessage)
  } catch (err) {
    console.error('Erreur envoi prompt :', err.message)
    messages.value.push({
      Role: 'assistant',
      Content: `Erreur : ${err.message}`,
    })
  } finally {
    isThinking.value = false
    scrollToBottom()
    }
}
</script>

<template>
  <div class="flex flex-col h-screen bg-neutral-900 text-white">
    <!-- Partie scrollable avec les messages -->
    <div ref="chatContainer" class="flex-1 overflow-y-auto px-6 py-4 w-full max-w-screen-xl mx-auto">
      <div v-if="loading" class="text-gray-500">Chargement...</div>
      <div v-else-if="error" class="text-red-500">Erreur : {{ error }}</div>
      <div v-else class="flex flex-col gap-4">
        <MessageBubble
          v-for="(msg, index) in messages"
          :key="index"
          :content="msg.Content"
          :role="msg.Role"
        />
      </div>
    </div>

    <!-- Input en bas -->
    <PromptInput @submit="handlePrompt" :loading="isThinking" />
  </div>
</template>
