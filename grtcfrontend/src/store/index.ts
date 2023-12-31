import { defineStore } from 'pinia'
import { getStorage, setStorage, removeStorage } from '@/utils/storage'
import { sessionStore } from './session'
export { sessionStore }
import { userStore } from './user'
export { userStore }
import { closeWs } from '@/utils/socket'
export const mainStore = defineStore('main', {
  state: () => {
    return {
      themeList: [
        {
          name: '亮色主题',
          class: 'light',
          id: 1,
        },
        {
          name: '暗色主题',
          class: 'dark',
          id: 1,
        },
      ],
      themeSelect: getStorage('theme') || 'light',
      token: getStorage('token') || '',
      userInfo: getStorage('userInfo', 'object') || {},
      pointType: getStorage('pointType', 'object') || {
        session: false,
        address: false,
      },
      isLogout: false,
      logoutInfo: getStorage('logoutInfo'),
      isPermission: false,
    }
  },
  actions: {
    setTheme(theme: string) {
      setStorage('theme', theme)
      this.themeSelect = theme
    },
    changPoint(key: string, type: boolean) {
      this.pointType[key] = type
      setStorage('pointType', this.pointType)
    },
    setToken(token: string) {
      setStorage('token', token)
      this.token = token
    },
    async setUserInfo(data: Object) {
      setStorage('userInfo', data)
      this.userInfo = data
    },
    setLogoutInfo(data: Object) {
      setStorage('logoutInfo', data)
      this.logoutInfo = data
    },
    clearMessage(){
      this.isLogout = false
    },
    changePermission(permission: boolean) {
      this.isPermission = permission
    },
    logOut(isLogout: boolean = false) {
      this.isLogout = isLogout
      closeWs()
      this.token = ''
      this.userInfo = {}
      removeStorage('token')
      removeStorage('userInfo')
      const sessionStores = sessionStore()
      sessionStores.init()
      const userStores = userStore()
      userStores.init()
    },
  },
})
