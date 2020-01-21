package com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.LayoutController


import android.annotation.SuppressLint
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.IntentFilter
import android.os.Bundle
import android.support.v4.app.Fragment
import android.support.v4.content.LocalBroadcastManager
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.BackendController.RetrofitHandler.MainHandler
import com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.BackendController.RetrofitHandler.ResponseObject.DeviceList
import com.blogspot.fikriprayoga.www.microcontrollerbasedsystem.R
import com.google.firebase.messaging.FirebaseMessaging
import kotlinx.android.synthetic.main.fragment_home.*
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory

// TODO: Rename parameter arguments, choose names that match
// the fragment initialization parameters, e.g. ARG_ITEM_NUMBER
private const val ARG_PARAM1 = "param1"
private const val ARG_PARAM2 = "param2"

/**
 * A simple [Fragment] subclass.
 *
 */
class HomeFragment : Fragment() {
    //lateinit var deviceList: List<DeviceList>
    lateinit var mCopyright: TextView

    override fun onCreateView(inflater: LayoutInflater, container: ViewGroup?,
                              savedInstanceState: Bundle?): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_home, container, false)
    }

    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)

        mCopyright = view!!.findViewById(R.id.textView_fragment_home_2)

        getDevices()

        mCopyright.setOnClickListener {
            getDevices()

        }

        FirebaseMessaging.getInstance().subscribeToTopic("update").addOnCompleteListener {
            Log.d("skripsiku", "Success to get the topic")

        }

    }

    fun getDevices() {
        Log.d("skripsiku", "buat masuk device")
        val retrofit = Retrofit.Builder()
                .baseUrl("http://13.229.84.45:8080/")
                .addConverterFactory(GsonConverterFactory.create())
                .build()

        val api = retrofit.create(MainHandler::class.java)
        val call = api.getDevices()

        call.enqueue(object : Callback<DeviceList> {
            override fun onFailure(call: Call<DeviceList>, t: Throwable) {
                Log.d("skripsiku1", t.message)
                Log.d("skripsiku2", t.cause.toString())
            }

            @SuppressLint("SetTextI18n")
            override fun onResponse(call: Call<DeviceList>, response: Response<DeviceList>) {
                val responseBody = response.body()
                if (responseBody != null) {
                    textView_fragment_home_2.text = "${responseBody.Temperature} Celcius"

                }

            }

        })

    }

    override fun onStart() {
        super.onStart()
        LocalBroadcastManager.getInstance(context!!).registerReceiver(mHandler, IntentFilter("www.fikriprayoga.blogspot.co.id_FCM-MESSAGE"))
    }

    override fun onPause() {
        super.onPause()
        LocalBroadcastManager.getInstance(context!!).unregisterReceiver(mHandler)

    }

    private val mHandler = object : BroadcastReceiver() {
        override fun onReceive(context: Context, intent: Intent) {
            getDevices()

        }
    }

}
