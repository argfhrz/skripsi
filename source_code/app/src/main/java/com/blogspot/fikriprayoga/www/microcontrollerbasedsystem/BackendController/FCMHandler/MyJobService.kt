package com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.BackendController.FCMHandler

import android.os.Bundle
import com.firebase.jobdispatcher.JobParameters
import com.firebase.jobdispatcher.JobService

class MyJobService : JobService() {
    private var myBundle: Bundle? = null

    override fun onStopJob(job: JobParameters?): Boolean {
        return false

    }

    override fun onStartJob(job: JobParameters?): Boolean {
        myBundle = job!!.extras

        return false

    }
}