const API_URL = "http://localhost:8080"

export async function fetchConversation(convId) {
  const response = await fetch(`${API_URL}/conv/${convId}`)

  if (!response.ok) {
    const errorData = await response.json()
    throw new Error(errorData.error || "Erreur lors du chargement de la conversation")
  }

  const messages = await response.json()
  return messages
}
