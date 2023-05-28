import { defineNuxtConfig } from 'nuxt/config'
import tailwindTypography from '@tailwindcss/typography'
import tailwindForms from '@tailwindcss/forms'

export default defineNuxtConfig({
  modules: ['@nuxtjs/apollo', '@nuxtjs/tailwindcss', 'nuxt-headlessui'],

  apollo: {
    clients: {
      default: {
        httpEndpoint: 'http://127.0.0.1:8000/graphql'
      }
    },
  },

  tailwindcss: {
    config: {
      plugins: [tailwindTypography, tailwindForms]
    }
  },
})
