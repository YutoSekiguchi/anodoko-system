export const state = () => ({
  // get
  scenesInfo: [],
  scenesVolList: [],
  scenePages: [],
  scenesUnderVolList: [],
  userSetScenesList: [],
  scene: [],
  userScenesInfo: [],

  // post
  addedSceneData: {},
})

export const mutations = {
  SET_SCENES_INFO(state, list) {
    state.scenesInfo = list
  },
  SET_SCENES_VOL_LIST(state, list) {
    state.scenesVolList = list
  },
  SET_SCENE_PAGES_ASC(state, list) {
    state.scenePages = list
  },
  SET_SCENE_PAGES_DESC(state, list) {
    state.scenePages = list
  },
  SET_SCENES_UNDER_VOL_LIST(state, list) {
    state.scenesUnderVolList = list
  },
  SET_USER_SET_SCENES_LIST(state, list) {
    state.userSetScenesList = list
  },
  SET_SCENE(state, list) {
    state.scene = list
  },
  SET_USER_SCENES_INFO(state, list) {
    state.userScenesInfo = list
  },
  SET_ADDED_SCENE_DATA(state, obj) {
    state.addedSceneData = obj
  },
  RESET_SCENE_DATA(state, n) {
    state.sceneInfo = null
    state.scenesInfo = n
    state.scenesVolList = n
    state.scenePages = n
    state.scene = n
  }
}

export const actions = {
  // 指定した漫画ページのシーンを取得
  async fetchScenesInfo({ commit }, [cid, volume, page]) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/info/${cid}/${volume}/${page}`,
    }).then((res) => {
      commit('SET_SCENES_INFO', res.data)
    }).catch(() => {})
  },
  // 指定した漫画の巻のシーン一覧
  async fetchScenesVolList({ commit }, [cid, volume]) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/volume/pages/${cid}/${volume}`
    }).then((res) => {
      commit('SET_SCENES_VOL_LIST', res.data)
    }).catch(() => {})
  },
  // 検索結果漫画のシーン昇順
  async fetchScenePagesAsc({ commit }, [chid1, chid2, chid3, chid4, chid5, cid, volume, scene]) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/pages/${chid1}/${chid2}/${chid3}/${chid4}/${chid5}/${cid}/${volume}/${scene}/asc`,
    }).then((res) => {
      commit('SET_SCENE_PAGES_ASC', res.data)
    }).catch(() => {})
  },
  // 検索結果漫画のシーン降順
  async fetchScenePagesDesc({ commit }, [chid1, chid2, chid3, chid4, chid5, cid, volume, scene]) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/pages/${chid1}/${chid2}/${chid3}/${chid4}/${chid5}/${cid}/${volume}/${scene}/desc`,
    }).then((res) => {
      commit('SET_SCENE_PAGES_DESC', res.data)
    }).catch(() => {})
  },
  // 検索候補の一覧
  async fetchScenesUnderVolList({ commit }, [chid1, chid2, chid3, chid4, chid5, cid, volume, scene]) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/volume/pages/under/${chid1}/${chid2}/${chid3}/${chid4}/${chid5}/${cid}/${volume}/${scene}`,
    }).then((res) => {
      commit('SET_SCENES_UNDER_VOL_LIST', res.data)
    }).catch(() => {})
  },
  //ユーザーが登録したシーン一覧
  async fetchUserSetSceneList({ commit }, uid) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/list/user/${uid}`,
    }).then((res) => {
      commit('SET_USER_SET_SCENES_LIST', res.data)
    }).catch(() => {})
  },
  // 指定したシーンの取得
  async fetchScene({ commit }, scid) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/scene/${scid}`,
    }).then((res) => {
      commit('SET_SCENE', res.data)
    }).catch(() => {})
  },
  // uid, cvid, page, sceneを指定したシーン一覧
  async fetchUserScenesInfo({ commit }, [uid, cvid, page, scene]) {
    await this.$axios({
      method: 'GET',
      url: `/scenes/user/info/${uid}/${cvid}/${page}/${scene}`,
    }).then((res) => {
      commit('SET_USER_SCENES_INFO', res.data)
    }).catch(() => {})
  },
  // post
  // シーンの追加
  async postScene({ commit, dispatch, rootState }, obj) {
    let statusCode = null
    await this.$axios({
      method: 'POST',
      url: '/scenes',
      data: obj,
    }).then((res) => {
      commit('SET_ADDED_SCENE_DATA', res.data)
    }).catch((e) => {
      statusCode = e.response.status
    })

    if (statusCode == 401) {
      await dispatch('postScene', obj)
    }
  },
  //リセット
  async resetScene({ commit }) {
    commit('RESET_SCENE_DATA', [])
  }
}