package com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.LayoutController

import android.annotation.SuppressLint
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.IntentFilter
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.os.Handler
import android.support.v4.app.Fragment
import android.support.v4.content.LocalBroadcastManager
import android.support.v7.widget.CardView
import android.view.View
import android.widget.TextView
import com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.R

class MainActivity : AppCompatActivity() {
    lateinit var mTitle: TextView
    lateinit var mLoading: CardView


    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)



        mTitle = findViewById(R.id.textView_activity_main_2)
        mLoading = findViewById(R.id.cardView_activity_main_3)

        val handler = Handler()
        handler.postDelayed(Runnable {
            mTitle.visibility = View.GONE
            mLoading.visibility = View.GONE

            changeFragment(HomeFragment(), "HomeFragment")

        }, 1000)

    }

    @SuppressLint("MissingSuperCall")
    override fun onSaveInstanceState(outState: Bundle) {
        //No call for super(). Bug on API Level > 11.
    }

    // method for change fragment
    fun changeFragment(newFragment: Fragment, tag: String) {
        val fragmentManager = supportFragmentManager
        val fragmentTransaction = fragmentManager.beginTransaction()
        fragmentTransaction.replace(R.id.frameLayout_activity_main_1, newFragment, tag)
        fragmentTransaction.commit()

    }

    fun changeFragment3(layout: Int, fragment: Fragment){
        supportFragmentManager
                .beginTransaction()
                .addToBackStack(null)
                .add(layout, fragment)
                .commitAllowingStateLoss()

    }
}
