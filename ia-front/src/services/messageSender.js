// src/services/chatService.js

const API_URL = 'http://localhost:8080'

/**
 * Envoie un message utilisateur au backend
 * @param {Object} payload - L'objet à envoyer
 * @param {number} payload.conversation_id - ID de la conversation
 * @param {string} payload.model - Nom du modèle (ex: "openrouter/horizon-beta")
 * @param {Object} payload.message - Message à envoyer (doit contenir `role` et `content`)
 * @returns {Promise<Object>} - La réponse JSON du serveur
 */
export async function sendMessageToChat({ conversation_id, model, message }) {
  const response = await fetch(`${API_URL}/chat`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      conversation_id,
      model,
      message,
    }),
  })

  if (!response.ok) {
    const errorData = await response.json()
    throw new Error(errorData.error || 'Erreur lors de l’envoi du message')
  }

  return await response.json()
}
