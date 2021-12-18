<template>
  <v-app-bar
      app
      color="white"
      dark
      height="80"
    >
      <NuxtLink :to="{ path: '/' }">
        <div class="d-flex align-center" light>
          <v-img
              class="ml-1"
              alt="Anodoko Icon"
              :src="logoImage"
              transition="scroll-x-transition"
              width="80"
              height="80"
            />
            <v-img
              class="ml-1"
              alt="Anodoko Title Icon"
              :src="titleImage"
              transition="scroll-x-reverse-transition"
              width="200"
              height="80"
            />
        </div>
      </NuxtLink>
      
      <v-spacer></v-spacer>


      <div v-if="this.loaded">
        <div v-if="!this.isLoggedIn">
          <v-btn target="_blank" @click="firebaseAuth()">
            <span class="font-weight-bold">ログイン</span>
          </v-btn>
        </div>
        <div v-else>
          <NuxtLink :to="{ path: '/user' }">
            <v-avatar>
              <v-img
                ripple
                class="ml-1"
                alt="User Icon"
                :src="this.userData.Image"
                transition="scale-transition"
                width="50"
                height="50"
              />`
            </v-avatar>
          </NuxtLink>
        </div>
      </div>
    </v-app-bar>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data() {
    return {
      logoImage: require("@/assets/anodoko.png"),
      loaded: true,
      titleImage: require("@/assets/anodoko_title.png"),
    }
  },
  methods: {
    ...mapActions('user', ['firebaseAuth', 'signOut']),
  },
  computed: {
    ...mapGetters('user', ['isLoggedIn']),
    ...mapState('user', ['userData']),
  },
  watch: {
    userData(val, newVal) {
      this.profileImage = val.Img
      this.loaded = true
    },
  },
  created() {

  },
  mounted() {
    if (this.isLoggedIn) {
      this.loaded = true
    }
  },
}

</script>