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