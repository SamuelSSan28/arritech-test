// Country codes mapping with flag emojis
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
  '598': { flag: '🇺🇾', name: 'Uruguay', code: 'UY' },
  '591': { flag: '🇧🇴', name: 'Bolivia', code: 'BO' },
  '593': { flag: '🇪🇨', name: 'Ecuador', code: 'EC' },
  
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

/**
 * Extract country code from phone number
 * @param {string} phone - Phone number with or without + or country code
 * @returns {string|null} - Country code or null if not found
 */
export function extractCountryCode(phone) {
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

/**
 * Get country info from phone number
 * @param {string} phone - Phone number
 * @returns {object|null} - Country info with flag, name, code or null
 */
export function getCountryInfo(phone) {
  const countryCode = extractCountryCode(phone)
  return countryCode ? countryCodes[countryCode] : null
}

/**
 * Get flag emoji from phone number
 * @param {string} phone - Phone number
 * @returns {string} - Flag emoji or 🌍 (world) if not found
 */
export function getFlagFromPhone(phone) {
  const countryInfo = getCountryInfo(phone)
  return countryInfo ? countryInfo.flag : '🌍'
}

/**
 * Get country name from phone number
 * @param {string} phone - Phone number
 * @returns {string} - Country name or 'Unknown' if not found
 */
export function getCountryNameFromPhone(phone) {
  const countryInfo = getCountryInfo(phone)
  return countryInfo ? countryInfo.name : 'Unknown'
}

/**
 * Format phone number with country flag
 * @param {string} phone - Phone number
 * @returns {string} - Formatted phone with flag
 */
export function formatPhoneWithFlag(phone) {
  if (!phone) return '—'
  
  const flag = getFlagFromPhone(phone)
  const cleanPhone = phone.replace(/[^\d]/g, '')
  
  // Simple formatting for common lengths
  if (cleanPhone.length === 11 && cleanPhone.startsWith('55')) {
    // Brazilian format: (11) 98765-4321
    const area = cleanPhone.slice(2, 4)
    const number = cleanPhone.slice(4)
    const formatted = `(${area}) ${number.slice(0, 5)}-${number.slice(5)}`
    return `${flag} ${formatted}`
  } else if (cleanPhone.length === 10) {
    // US/Canada format: (555) 123-4567
    const area = cleanPhone.slice(0, 3)
    const prefix = cleanPhone.slice(3, 6)
    const line = cleanPhone.slice(6)
    const formatted = `(${area}) ${prefix}-${line}`
    return `${flag} ${formatted}`
  } else {
    // Generic format
    return `${flag} ${phone}`
  }
}

/**
 * Get all available country codes
 * @returns {object} - All country codes with their info
 */
export function getAllCountryCodes() {
  return countryCodes
} 