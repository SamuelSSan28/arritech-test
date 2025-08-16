<template>
  <el-form
    ref="userForm"
    :model="formData"
    :rules="validationRules"
    label-position="top"
    @submit.prevent="handleSubmit"
    class="user-form"
    :validate-on-rule-change="false"
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
          :class="{ 'disabled-field': isEditing }"
        />
        <div class="form-help-text" v-if="isEditing">
          <el-icon><Lock /></el-icon>
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
          :disabled-date="isDateDisabled"
          @change="calculateAge"
          size="large"
          format="DD/MM/YYYY"
          value-format="YYYY-MM-DD"
          :class="{ 'disabled-field': isEditing }"
        />
        <div class="age-display" v-if="calculatedAge">
          {{ calculatedAge }} years old
        </div>
        <div class="form-help-text" v-if="isEditing">
          <el-icon><Lock /></el-icon>
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
  </el-form>
</template>

<script>
import { getFlagFromPhone } from '../../utils/countryFlags'
import { Lock } from '@element-plus/icons-vue'
import { ref } from 'vue'

export default {
  name: 'UserForm',
  components: {
    Lock
  },
  props: {
    formData: {
      type: Object,
      required: true
    },
    isEditing: {
      type: Boolean,
      default: false
    },
    loading: {
      type: Boolean,
      default: false
    },
    calculatedAge: {
      type: Number,
      default: 0
    },
    validationRules: {
      type: Object,
      required: true
    },
    errors: {
      type: Object,
      default: () => ({})
    }
  },
  emits: ['submit', 'calculate-age'],
  setup(props, { emit }) {
    const userForm = ref(null)

    const validate = (callback) => {
      if (!userForm.value) return
      userForm.value.validate(callback)
    }

    const handleSubmit = () => {
      validate((valid) => {
        if (valid) {
          emit('submit')
        }
      })
    }

    const calculateAge = () => {
      emit('calculate-age')
    }

    const isDateDisabled = (date) => {
      const today = new Date()
      const minDate = new Date()
      minDate.setFullYear(today.getFullYear() - 120)
      
      return date > today || date < minDate
    }

    return {  
      userForm,
      validate,
      handleSubmit,
      getFlagFromPhone,
      calculateAge,
      isDateDisabled
    }
  }
}
</script>

<style scoped>
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
  margin-top: 12px;
  font-size: 12px;
  color: #28a745;
  font-weight: 500;
  display: inline-block;
}

.form-help-text {
  margin-top: 12px;
  font-size: 12px;
  color: #6b7280;
  font-weight: 400;
  display: flex;
  align-items: center;
  gap: 4px;
}

.disabled-field {
  opacity: 0.7;
  background-color: #f8f9fa !important;
}

.disabled-field :deep(.el-input__wrapper) {
  background-color: #f8f9fa !important;
  border-color: #e5e7eb !important;
  color: #6b7280 !important;
}

.phone-input-container {
  position: relative;
  display: block;
  width: 100%;
}

.phone-input-container .el-input {
  width: 100%;
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

@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }
}
</style> 