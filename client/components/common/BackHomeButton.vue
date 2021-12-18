<template>
  <v-col align="center" class="mt-0 pt-0" cols="12">
    <v-btn class="py-5" color="error" elevation="13" @click="backToHome()" align="center"><v-icon dark left>mdi-home</v-icon>ホームへ戻る</v-btn>
  </v-col>
</template>

<script>
  export default {
    data: () => ({
  
    }),
    methods: {
      backToHome() {
        this.$router.push("/");
      }
    },
  }
</script>