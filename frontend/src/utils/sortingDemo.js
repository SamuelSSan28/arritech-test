// Demo file to showcase sorting functionality
import { userService } from '@/services/userService'

// Example sorting scenarios
export const sortingExamples = [
  {
    name: 'Sort by Name (A-Z)',
    params: { sortBy: 'name', sortDir: 'asc' },
    description: 'Users sorted alphabetically by name'
  },
  {
    name: 'Sort by Name (Z-A)',
    params: { sortBy: 'name', sortDir: 'desc' },
    description: 'Users sorted reverse alphabetically by name'
  },
  {
    name: 'Sort by Age (Oldest First)',
    params: { sortBy: 'age', sortDir: 'desc' },
    description: 'Users sorted by age, oldest first'
  },
  {
    name: 'Sort by Age (Youngest First)',
    params: { sortBy: 'age', sortDir: 'asc' },
    description: 'Users sorted by age, youngest first'
  },
  {
    name: 'Sort by Join Date (Newest First)',
    params: { sortBy: 'created_at', sortDir: 'desc' },
    description: 'Users sorted by join date, newest first'
  },
  {
    name: 'Sort by Join Date (Oldest First)',
    params: { sortBy: 'created_at', sortDir: 'asc' },
    description: 'Users sorted by join date, oldest first'
  }
]

// Function to demonstrate sorting
export async function demonstrateSorting() {
  console.log('🔄 Sorting Functionality Demo')
  console.log('==============================')
  
  for (const example of sortingExamples) {
    try {
      const response = await userService.getUsers({
        perPage: 3,
        ...example.params
      })
      
      const users = response.data.users
      const names = users.map(u => u.name).join(', ')
      
      console.log(`\n📊 ${example.name}`)
      console.log(`   ${example.description}`)
      console.log(`   First 3 users: ${names}`)
      
    } catch (error) {
      console.error(`Error with ${example.name}:`, error)
    }
  }
  
  console.log('\n✨ Features:')
  console.log('- Server-side sorting for performance')
  console.log('- Works with pagination')
  console.log('- Maintains search functionality')
  console.log('- Sortable columns: ID, Name, Age, Phone, Created At')
  console.log('- Sort directions: Ascending (asc) and Descending (desc)')
}

// Export for testing
export { sortingExamples } 