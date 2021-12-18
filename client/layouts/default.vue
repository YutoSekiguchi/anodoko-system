<template>
  <v-app dark>
    <Header />
    <v-main>
      <Nuxt />
    </v-main>
    <v-footer
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
import Header from '~/components/common/Header.vue'

export default {
  components: {Header},
}
</script>
