<template>
  <el-dialog
    v-model="dialogVisible"
    :title="isEditing ? 'Edit User' : 'Add New User'"
    width="600px"
    :before-close="handleClose"
    :close-on-click-modal="false"
    destroy-on-close
    align-center
    class="user-modal"
  >
    <div class="modal-content">
      <el-form
        ref="userForm"
        :model="formData"
        :rules="formRules"
        label-position="top"
        @submit.prevent="handleSubmit"
        class="user-form"
      >
        <div class="form-grid">
          <el-form-item label="Full Name" prop="name" required>
            <el-input
              v-model="formData.name"
              placeholder="Enter full name"
              maxlength="100"
              :disabled="loading"
              size="large"
            />
          </el-form-item>
          
          <el-form-item label="Email Address" prop="email" required>
            <el-input
              v-model="formData.email"
              type="email"
              placeholder="Enter email address"
              :disabled="loading || isEditing"
              size="large"
            />
            <div class="form-help-text" v-if="isEditing">
              Email cannot be changed after creation
            </div>
          </el-form-item>
          
          <el-form-item label="Date of Birth" prop="dateOfBirth" required>
            <el-date-picker
              v-model="formData.dateOfBirth"
              type="date"
              placeholder="Select date of birth"
              style="width: 100%"
              :disabled="loading || isEditing"
              :disabled-date="disabledDate"
              @change="calculateAge"
              size="large"
              format="DD/MM/YYYY"
              value-format="YYYY-MM-DD"
            />
            <div class="age-display" v-if="calculatedAge">
              {{ calculatedAge }} years old
            </div>
            <div class="form-help-text" v-if="isEditing">
              Date of birth cannot be changed after creation
            </div>
          </el-form-item>
          
          <el-form-item label="Phone Number" prop="phone">
            <div class="phone-input-container">
              <el-input
                v-model="formData.phone"
                placeholder="Enter phone number (optional)"
                maxlength="20"
                :disabled="loading"
                size="large"
              />
              <div class="phone-flag" v-if="formData.phone">
                {{ getFlagFromPhone(formData.phone) }}
              </div>
            </div>
          </el-form-item>
        </div>
        
        <el-form-item label="Address" prop="address" class="full-width">
          <el-input
            v-model="formData.address"
            type="textarea"
            :rows="3"
            placeholder="Enter complete address (optional)"
            maxlength="500"
            :disabled="loading"
            resize="none"
          />
        </el-form-item>
        
        <div class="form-actions">
          <el-button 
            @click="handleClose" 
            :disabled="loading" 
            size="large"
            class="cancel-btn"
          >
            Cancel
          </el-button>
          <el-button 
            type="primary" 
            @click="handleSubmit"
            :loading="loading"
            size="large"
            class="submit-btn"
          >
            {{ isEditing ? 'Update User' : 'Create User' }}
          </el-button>
        </div>
      </el-form>
    </div>
  </el-dialog>
</template>

<script>
import { ref, reactive, computed, watch, nextTick } from 'vue'
import { ElMessage } from 'element-plus'
import { userService } from '@/services/userService'
import { isValidEmail, isValidPhone } from '@/utils/helpers'
import { getFlagFromPhone } from '@/utils/countryFlags'

export default {
  name: 'UserModal',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    user: {
      type: Object,
      default: null
    }
  },
  emits: ['update:visible', 'user-saved'],
  setup(props, { emit }) {
    const userForm = ref(null)
    const loading = ref(false)
    const calculatedAge = ref(0)
    
    const formData = reactive({
      name: '',
      email: '',
      dateOfBirth: null,
      phone: '',
      address: ''
    })

    const dialogVisible = computed({
      get: () => props.visible,
      set: (value) => emit('update:visible', value)
    })

    const isEditing = computed(() => !!props.user?.id)

    const formRules = {
      name: [
        { required: true, message: 'Name is required', trigger: 'blur' },
        { min: 2, max: 100, message: 'Name must be 2-100 characters', trigger: 'blur' }
      ],
      email: [
        { required: true, message: 'Email is required', trigger: 'blur' },
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
          trigger: 'blur' 
        }
      ],
      dateOfBirth: [
        { required: true, message: 'Date of birth is required', trigger: 'change' },
        { 
          validator: (rule, value, callback) => {
            if (!value) {
              callback(new Error('Date of birth is required'))
            } else {
              const age = calculateAgeFromDate(value)
              if (age <= 18) {
                callback(new Error('User must be older than 18'))
              } else {
                callback()
              }
            }
          }, 
          trigger: 'change' 
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
          trigger: 'blur' 
        }
      ]
    }

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

    const disabledDate = (time) => {
      const today = new Date()
      const minDate = new Date()
      minDate.setFullYear(today.getFullYear() - 120)
      
      return time.getTime() > Date.now() || time.getTime() < minDate.getTime()
    }

    const loadUser = () => {
      if (props.user) {
        Object.assign(formData, {
          name: props.user.name || '',
          email: props.user.email || '',
          phone: props.user.phone || '',
          address: props.user.address || ''
        })
        
        // Use date_of_birth if available, otherwise calculate from age (for backward compatibility)
        if (props.user.date_of_birth) {
          formData.dateOfBirth = props.user.date_of_birth.split('T')[0] // Extract date part
          calculatedAge.value = calculateAgeFromDate(formData.dateOfBirth)
        } else if (props.user.age) {
          // Fallback: calculate approximate date from age
          const today = new Date()
          const birthYear = today.getFullYear() - props.user.age
          formData.dateOfBirth = `${birthYear}-${String(today.getMonth() + 1).padStart(2, '0')}-${String(today.getDate()).padStart(2, '0')}`
          calculatedAge.value = props.user.age
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
      
      if (userForm.value) {
        userForm.value.resetFields()
      }
    }

    const handleSubmit = async () => {
      if (!userForm.value) return
      
      try {
        await userForm.value.validate()
        loading.value = true
        
        const userData = {
          name: formData.name.trim(),
          email: formData.email.trim().toLowerCase(),
          date_of_birth: formData.dateOfBirth,
          phone: formData.phone.trim() || undefined,
          address: formData.address.trim() || undefined
        }

        if (isEditing.value) {
          // Remove email and date_of_birth for updates
          delete userData.email
          delete userData.date_of_birth
          await userService.updateUser(props.user.id, userData)
          ElMessage.success('User updated successfully!')
        } else {
          await userService.createUser(userData)
          ElMessage.success('User created successfully!')
        }
        
        emit('user-saved')
        handleClose()
      } catch (error) {
        if (error.errors) return
        console.error('Error saving user:', error)
      } finally {
        loading.value = false
      }
    }

    const handleClose = () => {
      dialogVisible.value = false
      setTimeout(resetForm, 300)
    }

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

    const getFlagFromPhone = (phone) => {
      const countryInfo = getCountryInfo(phone)
      return countryInfo ? countryInfo.flag : '🌍'
    }

    const getCountryInfo = (phone) => {
      const countryCode = extractCountryCode(phone)
      return countryCode ? countryCodes[countryCode] : null
    }

    const extractCountryCode = (phone) => {
      if (!phone) return null
      
      // Remove all non-digit characters except +
      const cleanPhone = phone.replace(/[^\d+]/g, '')
      
      // If starts with +, remove it
      const phoneWithoutPlus = cleanPhone.startsWith('+') ? cleanPhone.slice(1) : cleanPhone
      
      // Try to match country codes (longest first)
      const sortedCodes = Object.keys(countryCodes).sort((a, b) => b.length - a.length)
      
      for (const code of sortedCodes) {
        if (phoneWithoutPlus.startsWith(code)) {
          return code
        }
      }
      
      return null
    }

    const countryCodes = {
      // Brazil
      '55': { flag: '🇧🇷', name: 'Brazil', code: 'BR' },
      
      // North America
      '1': { flag: '🇺🇸', name: 'United States/Canada', code: 'US/CA' },
      
      // Europe
      '33': { flag: '🇫🇷', name: 'France', code: 'FR' },
      '44': { flag: '🇬🇧', name: 'United Kingdom', code: 'GB' },
      '49': { flag: '🇩🇪', name: 'Germany', code: 'DE' },
      '34': { flag: '🇪🇸', name: 'Spain', code: 'ES' },
      '39': { flag: '🇮🇹', name: 'Italy', code: 'IT' },
      '31': { flag: '🇳🇱', name: 'Netherlands', code: 'NL' },
      '32': { flag: '🇧🇪', name: 'Belgium', code: 'BE' },
      '41': { flag: '🇨🇭', name: 'Switzerland', code: 'CH' },
      '43': { flag: '🇦🇹', name: 'Austria', code: 'AT' },
      '46': { flag: '🇸🇪', name: 'Sweden', code: 'SE' },
      '47': { flag: '🇳🇴', name: 'Norway', code: 'NO' },
      '45': { flag: '🇩🇰', name: 'Denmark', code: 'DK' },
      '48': { flag: '🇵🇱', name: 'Poland', code: 'PL' },
      '420': { flag: '🇨🇿', name: 'Czech Republic', code: 'CZ' },
      '36': { flag: '🇭🇺', name: 'Hungary', code: 'HU' },
      '30': { flag: '🇬🇷', name: 'Greece', code: 'GR' },
      '351': { flag: '🇵🇹', name: 'Portugal', code: 'PT' },
      '380': { flag: '🇺🇦', name: 'Ukraine', code: 'UA' },
      '7': { flag: '🇷🇺', name: 'Russia', code: 'RU' },
      
      // Asia
      '86': { flag: '🇨🇳', name: 'China', code: 'CN' },
      '81': { flag: '🇯🇵', name: 'Japan', code: 'JP' },
      '82': { flag: '🇰🇷', name: 'South Korea', code: 'KR' },
      '91': { flag: '🇮🇳', name: 'India', code: 'IN' },
      '65': { flag: '🇸🇬', name: 'Singapore', code: 'SG' },
      '60': { flag: '🇲🇾', name: 'Malaysia', code: 'MY' },
      '66': { flag: '🇹🇭', name: 'Thailand', code: 'TH' },
      '84': { flag: '🇻🇳', name: 'Vietnam', code: 'VN' },
      '62': { flag: '🇮🇩', name: 'Indonesia', code: 'ID' },
      '63': { flag: '🇵🇭', name: 'Philippines', code: 'PH' },
      '852': { flag: '🇭🇰', name: 'Hong Kong', code: 'HK' },
      '886': { flag: '🇹🇼', name: 'Taiwan', code: 'TW' },
      
      // Oceania
      '61': { flag: '🇦🇺', name: 'Australia', code: 'AU' },
      '64': { flag: '🇳🇿', name: 'New Zealand', code: 'NZ' },
      
      // Africa
      '27': { flag: '🇿🇦', name: 'South Africa', code: 'ZA' },
      '20': { flag: '🇪🇬', name: 'Egypt', code: 'EG' },
      '234': { flag: '🇳🇬', name: 'Nigeria', code: 'NG' },
      '254': { flag: '🇰🇪', name: 'Kenya', code: 'KE' },
      '212': { flag: '🇲🇦', name: 'Morocco', code: 'MA' },
      
      // South America
      '54': { flag: '🇦🇷', name: 'Argentina', code: 'AR' },
      '56': { flag: '🇨🇱', name: 'Chile', code: 'CL' },
      '57': { flag: '🇨🇴', name: 'Colombia', code: 'CO' },
      '58': { flag: '🇻🇪', name: 'Venezuela', code: 'VE' },
      '51': { flag: '🇵🇪', name: 'Peru', code: 'PE' },
      '593': { flag: '🇪🇨', name: 'Ecuador', code: 'EC' },
      '595': { flag: '🇵🇾', name: 'Paraguay', code: 'PY' },
      '598': { flag: '🇺🇾', name: 'Uruguay', code: 'UY' },
      '591': { flag: '🇧🇴', name: 'Bolivia', code: 'BO' },
      
      // Middle East
      '972': { flag: '🇮🇱', name: 'Israel', code: 'IL' },
      '971': { flag: '🇦🇪', name: 'UAE', code: 'AE' },
      '966': { flag: '🇸🇦', name: 'Saudi Arabia', code: 'SA' },
      '90': { flag: '🇹🇷', name: 'Turkey', code: 'TR' },
      '98': { flag: '🇮🇷', name: 'Iran', code: 'IR' },
      '93': { flag: '🇦🇫', name: 'Afghanistan', code: 'AF' },
      
      // Central America & Caribbean
      '52': { flag: '🇲🇽', name: 'Mexico', code: 'MX' },
      '502': { flag: '🇬🇹', name: 'Guatemala', code: 'GT' },
      '503': { flag: '🇸🇻', name: 'El Salvador', code: 'SV' },
      '504': { flag: '🇭🇳', name: 'Honduras', code: 'HN' },
      '505': { flag: '🇳🇮', name: 'Nicaragua', code: 'NI' },
      '506': { flag: '🇨🇷', name: 'Costa Rica', code: 'CR' },
      '507': { flag: '🇵🇦', name: 'Panama', code: 'PA' },
      '1-809': { flag: '🇩🇴', name: 'Dominican Republic', code: 'DO' },
      '1-787': { flag: '🇵🇷', name: 'Puerto Rico', code: 'PR' },
      '1-876': { flag: '🇯🇲', name: 'Jamaica', code: 'JM' }
    }

    return {
      userForm,
      loading,
      calculatedAge,
      formData,
      formRules,
      dialogVisible,
      isEditing,
      calculateAge,
      disabledDate,
      handleSubmit,
      handleClose,
      getFlagFromPhone,
      getCountryInfo,
      extractCountryCode,
      countryCodes
    }
  }
}
</script>

<style scoped>
.user-modal :deep(.el-dialog) {
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.15);
}

.user-modal :deep(.el-dialog__header) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 28px 32px 20px 32px;
  margin: 0;
  border-bottom: none;
  position: relative;
}

.user-modal :deep(.el-dialog__header::after) {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #ff6b6b, #feca57, #48dbfb, #ff9ff3);
}

.user-modal :deep(.el-dialog__title) {
  color: white;
  font-size: 26px;
  font-weight: 700;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.user-modal :deep(.el-dialog__headerbtn) {
  top: 24px;
  right: 24px;
  width: 44px;
  height: 44px;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
  border: 1px solid rgba(255, 255, 255, 0.3);
}

.user-modal :deep(.el-dialog__headerbtn:hover) {
  background: rgba(255, 255, 255, 0.3);
  transform: scale(1.05);
}

.user-modal :deep(.el-dialog__close) {
  color: white;
  font-size: 22px;
  font-weight: bold;
}

.user-modal :deep(.el-dialog__body) {
  padding: 0;
}

.user-modal :deep(.el-dialog__footer) {
  display: none;
}

.modal-content {
  padding: 32px 32px 24px 32px;
}

.user-form {
  width: 100%;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  margin-bottom: 24px;
}

.full-width {
  grid-column: 1 / -1;
  margin-bottom: 28px;
}

.age-display {
  margin-top: 8px;
  font-size: 14px;
  color: #28a745;
  font-weight: 500;
  padding: 4px 8px;
  background: #f8fff9;
  border-radius: 6px;
  display: inline-block;
}

.form-help-text {
  margin-top: 8px;
  font-size: 12px;
  color: #6c757d;
  font-style: italic;
  padding: 4px 8px;
  background: #f8f9fa;
  border-radius: 6px;
  display: inline-block;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #e9ecef;
}

.cancel-btn {
  border-radius: 12px !important;
  background: #f8f9fa !important;
  border-color: #dee2e6 !important;
  color: #6c757d !important;
  font-weight: 500 !important;
  padding: 12px 24px !important;
  transition: all 0.3s ease !important;
}

.cancel-btn:hover {
  background: #e9ecef !important;
  border-color: #adb5bd !important;
  transform: translateY(-1px) !important;
}

.submit-btn {
  border-radius: 12px !important;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
  border: none !important;
  font-weight: 600 !important;
  padding: 12px 28px !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3) !important;
}

.submit-btn:hover {
  transform: translateY(-2px) !important;
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4) !important;
}

:deep(.el-form-item__label) {
  font-weight: 600;
  color: #495057;
  margin-bottom: 10px;
  font-size: 14px;
}

:deep(.el-form-item) {
  margin-bottom: 0;
}

:deep(.el-input__wrapper) {
  border-radius: 12px !important;
  border: 2px solid #e9ecef !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05) !important;
}

:deep(.el-input__wrapper:hover) {
  border-color: #667eea !important;
  box-shadow: 0 4px 8px rgba(102, 126, 234, 0.1) !important;
}

:deep(.el-input__wrapper.is-focus) {
  border-color: #667eea !important;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1) !important;
}

:deep(.el-input__wrapper.is-disabled) {
  background-color: #f8f9fa !important;
  border-color: #e9ecef !important;
}

:deep(.el-textarea__inner) {
  border-radius: 12px !important;
  border: 2px solid #e9ecef !important;
  transition: all 0.3s ease !important;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05) !important;
}

:deep(.el-textarea__inner:hover) {
  border-color: #667eea !important;
  box-shadow: 0 4px 8px rgba(102, 126, 234, 0.1) !important;
}

:deep(.el-textarea__inner:focus) {
  border-color: #667eea !important;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1) !important;
}

.phone-input-container {
  position: relative;
  display: flex;
  align-items: center;
}

.phone-flag {
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  font-size: 20px;
  pointer-events: none;
  z-index: 1;
  background: white;
  padding: 4px;
  border-radius: 6px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

@media (max-width: 768px) {
  .user-modal :deep(.el-dialog) {
    width: 95% !important;
    margin: 4% auto !important;
  }
  
  .modal-content {
    padding: 24px 24px 20px 24px;
  }
  
  .form-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }
  
  .form-actions {
    flex-direction: column;
    gap: 12px;
  }
  
  .submit-btn, .cancel-btn {
    width: 100% !important;
  }
}
</style> 