<template>
  <v-container >
    <v-row>
    <v-col cols="6" sm="4" class="pr-0">
      <v-navigation-drawer permanent width="450px">
      <v-system-bar color="orange"></v-system-bar>
      <v-list>
        <v-list-item>
          <v-list-item-avatar>
            <v-img
                class=""
                alt="User Icon"
                :src="this.$store.state.user.userData.Image"
                transition="scale-transition"
                width="70"
                height="70"
              />
          </v-list-item-avatar>
        </v-list-item>

        <v-list-item link>
          <v-list-item-content >
            <v-list-item-title class="text-h6">
              {{this.$store.state.user.userData.Name}}
            </v-list-item-title>
            <v-list-item-subtitle>{{this.$store.state.user.userData.Mail}}</v-list-item-subtitle>
          </v-list-item-content>

          <v-list-item-action>
            <v-icon>mdi-menu-down</v-icon>
          </v-list-item-action>
        </v-list-item>
      </v-list>
      <v-divider></v-divider>
      <v-list
        nav
        dense
      >
        <v-list-item-group
          v-model="selectedItem"
          color="primary"
        >
          <v-list-item
            v-for="(item, i) in items"
            :key="i"
          >
            <v-list-item-icon>
              <v-icon v-text="item.icon"></v-icon>
            </v-list-item-icon>
              
            <v-list-item-content>
              <v-list-item-title v-text="item.text"></v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
    </v-col>
    
      <v-col cols="6" sm="8" style="overflow: scroll;" class="pl-0">
        <v-card
          class="mx-auto"
          max-width="1000"
          min-height="410"
          max-height="410"
          tile
          style="overflow: scroll;"
        >
        <div v-if="this.selectedItem==0" class="pl-0">
          <FavoriteScene />
        </div>

        <div v-if="this.selectedItem==1" >
          <RegisterScenes />
        </div>
        </v-card>
      </v-col>
    
    </v-row>
  </v-container>
</template>

<script>
import FavoriteScene from '~/components/user/FavoriteScene.vue'
import RegisterScenes from '~/components/user/RegisterScenes.vue'
// import { mapState, mapActions } from 'vuex'

export default {
  components: {
    FavoriteScene,
    RegisterScenes,
  },
  data() {
    return {
      selectedItem: null,
      items: [
        { text: 'お気に入り', icon: 'mdi-heart' },
        { text: 'あなたの登録シーン', icon: 'mdi-checkbox-blank-circle' },
        { text: '閲覧履歴？', icon: 'mdi-history' },
        { text: 'その他？', icon: 'mdi-checkbox-blank-circle' },
        { text: 'その他2？', icon: 'mdi-checkbox-blank-circle' },
      ],
    }
  },
  
}
</script>