// Demo file to showcase country flag functionality
import { formatPhoneWithFlag, getFlagFromPhone, getCountryNameFromPhone } from './countryFlags'

// Example phone numbers from different countries
const demoPhones = [
  '+5511987654321',    // 🇧🇷 Brazil
  '+15551234567',      // 🇺🇸 United States
  '+447911123456',     // 🇬🇧 United Kingdom
  '+33123456789',      // 🇫🇷 France
  '+49123456789',      // 🇩🇪 Germany
  '+81901234567',      // 🇯🇵 Japan
  '+86123456789',      // 🇨🇳 China
  '+34612345678',      // 🇪🇸 Spain
  '+39391234567',      // 🇮🇹 Italy
  '+61412345678',      // 🇦🇺 Australia
  '+27123456789',      // 🇿🇦 South Africa
  '+54123456789',      // 🇦🇷 Argentina
  '+52551234567',      // 🇲🇽 Mexico
  '+972123456789',     // 🇮🇱 Israel
  '+971123456789'      // 🇦🇪 UAE
]

// Function to demonstrate the functionality
export function demonstrateCountryFlags() {
  console.log('🌍 Country Flag Phone Number Demo')
  console.log('=====================================')
  
  demoPhones.forEach(phone => {
    const flag = getFlagFromPhone(phone)
    const country = getCountryNameFromPhone(phone)
    const formatted = formatPhoneWithFlag(phone)
    
    console.log(`${flag} ${country}: ${formatted}`)
  })
  
  console.log('\n✨ Features:')
  console.log('- Automatic country detection from phone number')
  console.log('- Flag emoji display')
  console.log('- Country name identification')
  console.log('- Phone number formatting')
  console.log('- Support for 50+ countries')
}

// Export demo phones for testing
export { demoPhones } 