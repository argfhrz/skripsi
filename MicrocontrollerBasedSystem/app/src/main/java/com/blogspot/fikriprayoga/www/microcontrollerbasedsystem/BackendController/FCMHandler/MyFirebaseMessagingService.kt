package com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.BackendController.FCMHandler

import android.content.Intent
import android.os.Bundle
import android.util.Log
import com.google.firebase.messaging.FirebaseMessagingService
import com.google.firebase.messaging.RemoteMessage
import android.support.design.widget.CoordinatorLayout.Behavior.setTag
import android.support.v4.content.LocalBroadcastManager
import android.widget.Toast
import com.firebase.jobdispatcher.FirebaseJobDispatcher
import com.firebase.jobdispatcher.GooglePlayDriver


class MyFirebaseMessagingService : FirebaseMessagingService() {
    private val TAG = "MBS"

    override fun onNewToken(token: String?) {
        Log.d(TAG, "Refreshed token: " + token!!)

        sendRegistrationToServer(token)
    }

    private fun sendRegistrationToServer(token: String) {

    }

    override fun onMessageReceived(remoteMessage: RemoteMessage?) {
        Log.d(TAG, "From: " + remoteMessage!!.from!!)

        if (remoteMessage.data.size > 0) {
            Log.d(TAG, "Message data payload: " + remoteMessage.data)

            if (false) {
                scheduleJob()

            } else {
                handleNow()

            }

        }

        if (remoteMessage.notification != null) {
            Log.d(TAG, "Message Notification Body: " + remoteMessage.notification!!.body!!)
        }

    }

    private fun scheduleJob() {
        val dispatcher = FirebaseJobDispatcher(GooglePlayDriver(this))
        val myJob = dispatcher.newJobBuilder()
                .setService(MyJobService::class.java)
                .setTag("my-job-tag")
                .build()
        dispatcher.schedule(myJob)
    }

    private fun handleNow() {
        val intent = Intent("www.fikriprayoga.blogspot.co.id_FCM-MESSAGE")
        val localBroadcastManager = LocalBroadcastManager.getInstance(this)
        localBroadcastManager.sendBroadcast(intent)
        Log.d(TAG, "Short lived task is done.")
    }

}