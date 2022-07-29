import React, { createContext, useContext, useEffect, useState } from 'react'
import getAssets, { DataAssets } from '../services/assets'
import getProfile, { DataProfile } from '../services/profile'
import getWallets, { DataWallet } from '../services/wallets'

type AuthProviderWithFunc = {
  isLogin: boolean
  userProfile?: DataProfile
  assets?: DataAssets[]
  wallets?: DataWallet[]
  refresh: () => void
}
const initialState: AuthProviderWithFunc = {
  isLogin: false,
  userProfile: undefined,
  assets: [],
  wallets: [],
  refresh: () => null,
}
const AuthContext = createContext<AuthProviderWithFunc>(initialState)
export const useProfileAndAssets = (): AuthProviderWithFunc =>
  useContext(AuthContext)

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isLogin, setIsLogin] = useState(false)
  const [userProfile, setUserProfile] = useState<DataProfile>()
  const [assets, setAssets] = useState<DataAssets[]>([])
  const [wallets, setWallet] = useState<DataWallet[]>([])

  const handleGetData = async () => {
    const [profile, asset, wallet] = await Promise.all([
      await getProfile(),
      await getAssets(),
      await getWallets(),
    ])

    if (profile) {
      setUserProfile(profile.data)
      setIsLogin(true)
    }
    if (asset) {
      setAssets(asset.data || [])
    }
    if (wallet) {
      setWallet(wallet.data || [])
    }
  }

  useEffect(() => {
    handleGetData()
  }, [])

  return (
    <AuthContext.Provider
      value={{ isLogin, userProfile, assets, wallets, refresh: handleGetData }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider
