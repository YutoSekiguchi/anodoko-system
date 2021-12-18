import { initializeApp } from "firebase/app"
import { getAuth, onAuthStateChanged, onIdTokenChanged } from 'firebase/auth'
import config from '~/firebase.config'

initializeApp(config)
const auth = getAuth()

export default (context) => {
  const { store } = context
  return new Promise((resolve) => {
    onAuthStateChanged(auth, (user) => {
      if (user) {
        const displayName = user.displayName
        const email = user.email
        const photoURL = user.photoURL
        user.getIdToken().then((idToken) => {
          store.dispatch('user/setJwt', idToken.toString())
        })
        if (displayName && email && photoURL) {
          store.dispatch('user/signIn', {
            displayName,
            email,
            photoURL,
          })
        }
      } else {
        store.dispatch('user/signOut', {})
      }
      resolve()
    })

    // onIdTokenChanged(auth, (user) => {
    //   if (user) {
    //     user.getIdToken().then((idToken) => {
    //       store.dispatch('user/setJwt', idToken.toString())
    //     })
    //   } else {
    //     store.dispatch('user/signOut', {})
    //   }
    //   resolve()
    // })
  })
}