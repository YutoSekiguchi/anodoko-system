<template>
  <div class="button " align="center">
    <v-btn class="" color="error" @click="backPage($route.params.cid)"><v-icon dark left>mdi-arrow-left</v-icon>巻の選択に戻る</v-btn>
  </div>
</template>

<script>
export default {
  methods: {
    backPage(cid) {
      this.$router.push(`/search/${cid}`);
    }
  }
}
</script>