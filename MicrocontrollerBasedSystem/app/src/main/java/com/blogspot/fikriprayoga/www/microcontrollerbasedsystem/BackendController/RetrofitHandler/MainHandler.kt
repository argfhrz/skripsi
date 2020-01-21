package com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.BackendController.RetrofitHandler

import com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.BackendController.RetrofitHandler.ResponseObject.DeviceList
import retrofit2.Call
import retrofit2.http.GET

interface MainHandler {
    @GET("readData")
    fun getDevices(): Call<DeviceList>

}