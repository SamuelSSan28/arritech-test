import { ref, reactive, computed, watch } from 'vue'
import { isValidEmail, isValidPhone } from '../utils/helpers'

export function useUserForm(user = null, isEditing = false) {
  const formData = reactive({
    name: '',
    email: '',
    dateOfBirth: null,
    phone: '',
    address: ''
  })

  const calculatedAge = ref(0)
  const errors = ref({})

  // Business Rules
  const businessRules = {
    minAge: 18,
    maxAge: 120,
    nameMinLength: 2,
    nameMaxLength: 100,
    phoneMaxLength: 20,
    addressMaxLength: 500
  }

  // Validation Rules
  const validationRules = {
    name: [
      { required: true, message: 'Name is required', trigger: ['blur', 'submit'] },
      { 
        min: businessRules.nameMinLength, 
        max: businessRules.nameMaxLength, 
        message: `Name must be ${businessRules.nameMinLength}-${businessRules.nameMaxLength} characters`, 
        trigger: ['blur', 'submit']
      }
    ],
    email: [
      { required: true, message: 'Email is required', trigger: ['blur', 'submit'] },
      { 
        validator: (rule, value, callback) => {
          if (!value) {
            callback(new Error('Email is required'))
          } else if (!isValidEmail(value)) {
            callback(new Error('Please enter a valid email'))
          } else {
            callback()
          }
        }, 
        trigger: ['blur', 'submit']
      }
    ],
    dateOfBirth: [
      { required: true, message: 'Date of birth is required', trigger: ['change', 'submit'] },
      { 
        validator: (rule, value, callback) => {
          if (!value) {
            callback(new Error('Date of birth is required'))
          } else {
            const age = calculateAgeFromDate(value)
            if (age < businessRules.minAge) {
              callback(new Error(`User must be at least ${businessRules.minAge} years old`))
            } else if (age > businessRules.maxAge) {
              callback(new Error(`User cannot be older than ${businessRules.maxAge} years`))
            } else {
              callback()
            }
          }
        }, 
        trigger: ['change', 'submit']
      }
    ],
    phone: [
      { 
        validator: (rule, value, callback) => {
          if (value && !isValidPhone(value)) {
            callback(new Error('Please enter a valid phone number'))
          } else {
            callback()
          }
        }, 
        trigger: ['blur', 'submit']
      }
    ]
  }

  // Business Logic Functions
  const calculateAgeFromDate = (birthDate) => {
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

  const calculateAge = () => {
    calculatedAge.value = calculateAgeFromDate(formData.dateOfBirth)
  }

  const isDateDisabled = (time) => {
    const today = new Date()
    const minDate = new Date()
    minDate.setFullYear(today.getFullYear() - businessRules.maxAge)
    
    return time.getTime() > Date.now() || time.getTime() < minDate.getTime()
  }

  const validateForm = () => {
    const newErrors = {}
    
    // Name validation
    if (!formData.name.trim()) {
      newErrors.name = 'Name is required'
    } else if (formData.name.trim().length < businessRules.nameMinLength) {
      newErrors.name = `Name must be at least ${businessRules.nameMinLength} characters`
    } else if (formData.name.trim().length > businessRules.nameMaxLength) {
      newErrors.name = `Name cannot exceed ${businessRules.nameMaxLength} characters`
    }

    // Email validation
    if (!formData.email.trim()) {
      newErrors.email = 'Email is required'
    } else if (!isValidEmail(formData.email.trim())) {
      newErrors.email = 'Please enter a valid email'
    }

    // Date of birth validation
    if (!formData.dateOfBirth) {
      newErrors.dateOfBirth = 'Date of birth is required'
    } else {
      const age = calculateAgeFromDate(formData.dateOfBirth)
      if (age < businessRules.minAge) {
        newErrors.dateOfBirth = `User must be at least ${businessRules.minAge} years old`
      } else if (age > businessRules.maxAge) {
        newErrors.dateOfBirth = `User cannot be older than ${businessRules.maxAge} years`
      }
    }

    // Phone validation (optional)
    if (formData.phone.trim() && !isValidPhone(formData.phone.trim())) {
      newErrors.phone = 'Please enter a valid phone number'
    }

    // Address validation (optional)
    if (formData.address.trim() && formData.address.trim().length > businessRules.addressMaxLength) {
      newErrors.address = `Address cannot exceed ${businessRules.addressMaxLength} characters`
    }

    errors.value = newErrors
    return Object.keys(newErrors).length === 0
  }

  const loadUserData = (userData) => {
    if (userData) {
      Object.assign(formData, {
        name: userData.name || '',
        email: userData.email || '',
        phone: userData.phone || '',
        address: userData.address || ''
      })
      
      if (userData.date_of_birth) {
        formData.dateOfBirth = userData.date_of_birth.split('T')[0]
        calculatedAge.value = calculateAgeFromDate(formData.dateOfBirth)
      } else if (userData.age) {
        const today = new Date()
        const birthYear = today.getFullYear() - userData.age
        formData.dateOfBirth = `${birthYear}-${String(today.getMonth() + 1).padStart(2, '0')}-${String(today.getDate()).padStart(2, '0')}`
        calculatedAge.value = userData.age
      }
    }
  }

  const resetForm = () => {
    Object.assign(formData, {
      name: '',
      email: '',
      dateOfBirth: null,
      phone: '',
      address: ''
    })
    calculatedAge.value = 0
    errors.value = {}
  }

  const getFormData = () => ({
    name: formData.name.trim(),
    email: formData.email.trim().toLowerCase(),
    date_of_birth: formData.dateOfBirth,
    phone: formData.phone.trim() || undefined,
    address: formData.address.trim() || undefined
  })

  const getUpdateData = () => {
    const data = getFormData()
    // Remove immutable fields for updates
    delete data.email
    delete data.date_of_birth
    return data
  }

  // Watchers
  watch(() => formData.dateOfBirth, () => {
    calculateAge()
  })

  // Initialize if user is provided
  if (user) {
    loadUserData(user)
  }

  return {
    // State
    formData,
    calculatedAge,
    errors,
    
    // Business Rules
    businessRules,
    
    // Validation
    validationRules,
    validateForm,
    
    // Business Logic
    calculateAgeFromDate,
    calculateAge,
    isDateDisabled,
    
    // Form Operations
    loadUserData,
    resetForm,
    getFormData,
    getUpdateData
  }
} 