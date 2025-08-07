<script setup>
import MessageBubble from './components/MessageBubble.vue'
import { ref } from 'vue'
import Conversation from './components/Conversation.vue'

</script>

<template>
  <div class="min-h-screen bg-neutral-900 text-white">
    <main class="max-w-screen-xl mx-auto p-4">
      <Conversation convId="1" />
    </main>
  </div>
</template>
