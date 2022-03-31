const {PHASE_DEVELOPMENT_SERVER} = require('next/constants') 

module.exports = (phase) => {
  if (phase === PHASE_DEVELOPMENT_SERVER) {
    return{
      env: {
        NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL,
        NEXTAUTH_URL: process.env.NEXTAUTH_URL,
      },
      images: {
        domains: ['images.unsplash.com', 'frontend.myapp.com'],
      }
    }
  }
  return {
    env: {
      NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL,
      NEXTGATEWAY_PORT: process.env.NEXTGATEWAY_PORT,
    },
    images: {
      domains: ['images.unsplash.com', 'frontend.myapp.com/'],
    }
  }
}
