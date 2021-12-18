<template>
<div>
  <div v-for="(data, i) in this.userSetScenesList" :key="i">
    <v-row style="border-bottom: 0.5px solid grey;" class="mr-12 mt-3">
      <Scenes :data="data" />
      <ScenesDetail :data="data" />
    </v-row>
  </div>
</div>
</template>

<script>
import Scenes from '~/components/user/Scenes.vue'
import ScenesDetail from '~/components/user/ScenesDetail.vue'
import { mapState, mapActions } from 'vuex'

export default {
  components: {
    Scenes,
    ScenesDetail
  },
  methods: {
    ...mapActions('scene', [
      'fetchUserSetSceneList',
    ]),
  },
  computed: {
    ...mapState('scene', [
      'userSetScenesList',
    ])
  },
  async created() {
    await this.fetchUserSetSceneList(this.$store.state.user.userData.ID)
  }
}
</script>