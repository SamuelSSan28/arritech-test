// Format phone number
export function formatPhone(phone) {
  if (!phone) return '-'
  
  // Simple phone formatting - can be enhanced based on requirements
  const cleaned = phone.replace(/\D/g, '')
  const match = cleaned.match(/^(\d{3})(\d{3})(\d{4})$/)
  
  if (match) {
    return `(${match[1]}) ${match[2]}-${match[3]}`
  }
  
  return phone
}

// Validate phone format
export function isValidPhone(phone) {
  if (!phone) return true // Phone is optional
  const phoneRegex = /^[\+]?[1-9][\d]{0,15}$/
  return phoneRegex.test(phone.replace(/\s/g, ''))
} 