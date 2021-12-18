<template>
  <v-col cols="5" align="center">
    <h4 class="font-weight-bold">{{ data.Name }}</h4>
    <h5>{{ data.Volume }}巻{{ data.Page }}ページ</h5>
  </v-col>
</template>

<script>

export default {
  props: {
    data: {
      type: Object
    }
  }
}
</script>