// Debounce function to limit API calls
export function debounce(func, delay) {
  let timeoutId
  return function (...args) {
    clearTimeout(timeoutId)
    timeoutId = setTimeout(() => func.apply(this, args), delay)
  }
}

// Format date to readable string
export function formatDate(dateString, options = {}) {
  if (!dateString) return '-'
  
  const defaultOptions = {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    ...options
  }
  
  return new Date(dateString).toLocaleDateString('en-US', defaultOptions)
}

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

// Calculate age from date of birth
export function calculateAge(birthDate) {
  if (!birthDate) return 0
  
  const today = new Date()
  const birth = new Date(birthDate)
  let age = today.getFullYear() - birth.getFullYear()
  const monthDiff = today.getMonth() - birth.getMonth()
  
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
    age--
  }
  
  return age
}

// Calculate date of birth from age (approximate)
export function calculateDateOfBirth(age) {
  if (!age || age <= 0) return null
  
  const today = new Date()
  const birthYear = today.getFullYear() - age
  return new Date(birthYear, today.getMonth(), today.getDate())
}

// Format date of birth for display
export function formatDateOfBirth(dateString) {
  if (!dateString) return '-'
  
  const date = new Date(dateString)
  const age = calculateAge(date)
  
  return `${formatDate(dateString)} (${age} years old)`
}

// Validate email format
export function isValidEmail(email) {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(email)
}

// Validate age (must be greater than 18)
export function isValidAge(age) {
  return Number.isInteger(age) && age > 18 && age <= 120
}

// Validate date of birth (person must be older than 18)
export function isValidDateOfBirth(birthDate) {
  if (!birthDate) return false
  
  const age = calculateAge(birthDate)
  return isValidAge(age)
}

// Validate phone format
export function isValidPhone(phone) {
  if (!phone) return true // Phone is optional
  const phoneRegex = /^[\+]?[1-9][\d]{0,15}$/
  return phoneRegex.test(phone.replace(/\s/g, ''))
} 