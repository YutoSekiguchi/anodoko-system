<template>
  <v-col cols="6" align="start">
    <h5 class="orange--text pt-5">{{ data.DetailScene }}</h5>
  </v-col>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  props: {
    data: {
      type: Object
    }
  },
}
</script>