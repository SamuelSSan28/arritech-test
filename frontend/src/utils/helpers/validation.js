import { calculateAge } from './dateUtils.js'

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