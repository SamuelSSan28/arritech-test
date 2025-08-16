import { ref, computed, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { userService } from '../services/userService'
import { useUserForm } from './useUserForm'

export function useUserModal(props, emit) {
  const loading = ref(false)
  const userForm = ref(null)
  
  // Use the form composable
  const {
    formData,
    calculatedAge,
    errors,
    validationRules,
    validateForm,
    isDateDisabled,
    loadUserData,
    resetForm,
    getFormData,
    getUpdateData
  } = useUserForm(props.user, computed(() => !!props.user?.id))

  const dialogVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
  })

  const isEditing = computed(() => !!props.user?.id)

  // Modal Operations
  const handleSubmit = async () => {
    try {
      // Use business logic validation
      if (!validateForm()) {
        // Show first error
        const firstError = Object.values(errors.value)[0]
        ElMessage.error(firstError)
        return
      }

      loading.value = true
      
      if (isEditing.value) {
        const userData = getUpdateData()
        await userService.updateUser(props.user.id, userData)
        ElMessage.success('User updated successfully!')
      } else {
        const userData = getFormData()
        await userService.createUser(userData)
        ElMessage.success('User created successfully!')
      }
      
      emit('user-saved')
      handleClose()
    } catch (error) {
      if (error.errors) return
      console.error('Error saving user:', error)
      ElMessage.error('An error occurred while saving the user')
    } finally {
      loading.value = false
    }
  }

  const handleClose = () => {
    dialogVisible.value = false
    setTimeout(resetForm, 300)
  }

  // Lifecycle
  const loadUser = () => {
    if (props.user) {
      loadUserData(props.user)
    }
  }

  // Watchers
  watch(() => props.visible, (newVal) => {
    if (newVal) {
      nextTick(loadUser)
    }
  })

  watch(() => props.user, () => {
    if (props.visible) {
      loadUser()
    }
  })

  return {
    // State
    loading,
    userForm,
    formData,
    calculatedAge,
    errors,
    
    // Computed
    dialogVisible,
    isEditing,
    
    // Validation
    validationRules,
    
    // Business Logic
    isDateDisabled,
    
    // Methods
    handleSubmit,
    handleClose,
    resetForm
  }
} 