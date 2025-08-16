import api from './api'

export const userService = {
  // Get all users with pagination, search and sorting
  async getUsers(params = {}) {
    const queryParams = new URLSearchParams({
      page: params.page || 1,
      per_page: params.perPage || 10,
      ...(params.search && { search: params.search }),
      ...(params.sortBy && { sortBy: params.sortBy }),
      ...(params.sortDir && { sortDir: params.sortDir })
    })

    return await api.get(`/users?${queryParams}`)
  },

  // Get user by ID
  async getUser(id) {
    return await api.get(`/users/${id}`)
  },

  // Create new user
  async createUser(userData) {
    return await api.post('/users', userData)
  },

  // Update user
  async updateUser(id, userData) {
    return await api.put(`/users/${id}`, userData)
  },

  // Delete user
  async deleteUser(id) {
    return await api.delete(`/users/${id}`)
  }
} 