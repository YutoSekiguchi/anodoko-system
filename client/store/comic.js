export const state = () => ({
  // get
  comics: [],
  comic: [],
  comicVol: [],
  comicFirstVols: [],
  comicVols: [],
  comicUnderVols: [],

  // post
  addedComicData: {},
  addedComicVolData: {},
})

export const mutations = {
  SET_COMICS(state, list) {
    state.comics = list
  },
  SET_COMIC(state, list) {
    state.comic = list
  },
  SET_COMIC_VOL(state, list) {
    state.comicVol = list
  },
  SET_COMIC_FIRST_VOLS(state, list) {
    state.comicFirstVols = list
  },
  SET_COMIC_VOLS(state, list) {
    state.comicVols = list
  },
  SET_COMIC_UNDER_VOLS(state, list) {
    state.comicUnderVols = list
  },
  SET_ADDED_COMIC_DATA(state, obj) {
    state.addedComicData = obj
  },
  SET_ADDED_COMIC_VOL_DATA(state, obj) {
    state.addedComicVolData = obj
  },
  RESET_COMIC_DATA(state, n) {
    state.comics = n
    state.comic = n
    state.comicVol = n
    state.comicFirstVols = n
    state.comicVols = n
    state.comicUnderVols = n
  }
}

export const actions = {
  // 漫画一覧の取得
  async fetchComics({ commit }) {
    await this.$axios({
      method: 'GET',
      url: '/comics',
    }).then((res) => {
      commit('SET_COMICS', res.data)
    }).catch(() => {})
  },
  // 登録されてる漫画の一番巻数の小さい巻を取得
  async fetchComicFirstVols({ commit }) {
    await this.$axios({
      method: 'GET',
      url: '/comics/first',
    }).then((res) => {
      commit('SET_COMIC_FIRST_VOLS', res.data)
    }).catch(() => {})
  },
  // idを指定して漫画の取得
  async fetchComic({ commit }, id) {
    await this.$axios({
      method: 'GET',
      url: `/comics/${id}`,
    }).then((res) => {
      commit('SET_COMIC', res.data)
    }).catch(() => {})
  },
  // idを指定して漫画の巻情報を取得
  async fetchComicVol({ commit }, [cid, volume]) {
    await this.$axios({
      method: 'GET',
      url: `/comics/volume/${cid}/${volume}`,
    }).then((res) => {
      commit('SET_COMIC_VOL', res.data)
    }).catch(() => {})
  },
  // cidを指定して漫画の登録されている巻の取得
  async fetchComicVols({ commit }, cid) {
    await this.$axios({
      method: 'GET',
      url: `/comics/volumes/${cid}`,
    }).then((res) => {
      commit('SET_COMIC_VOLS', res.data)
    }).catch(() => {})
  },
  // cidとvolumeを指定してvolume以下のcidの漫画一覧を取得
  async fetchComicUnderVols({ commit }, [cid, volume]) {
    await this.$axios({
      method: 'GET',
      url: `/comics/under/${cid}/${volume}`,
    }).then((res) => {
      commit('SET_COMIC_UNDER_VOLS', res.data)
    }).catch(() => {})
  },
  // POST
  // 漫画を追加
  async postComic({ commit, dispatch, rootState }, obj) {
    let statusCode = null
    await this.$axios({
      method: 'POST',
      url: '/comics',
      data: obj,
    }).then((res) => {
      commit('SET_ADDED_COMIC_DATA', res.data)
    }).catch((e) => {
      console.log(e)
      statusCode = e.response.status
    })
    
    if (statusCode == 401) {
      await dispatch('postComic', obj)
    }
  },
  // 巻の追加
  async postComicVol({ commit, dispatch, rootState }, obj) {
    let statusCode = null
    await this.$axios({
      method: 'POST',
      url: '/comics/volumes',
      data: obj,
    }).then((res) => {
      commit('SET_ADDED_COMIC_VOL_DATA', res.data)
    }).catch((e) => {
      console.log(e)
      statusCode = e.response.status
    })

    if (statusCode == 401) {
      await dispatch('postComicVol', obj)
    }
  },

  //リセット
  async resetComic({ commit }) {
    commit('RESET_COMIC_DATA', [])
  }
}