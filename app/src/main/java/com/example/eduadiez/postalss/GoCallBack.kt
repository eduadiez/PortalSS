package com.example.eduadiez.postalss

import gomobile.JavaCallback
import android.app.Activity
import kotlinx.android.synthetic.main.activity_main.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch

class GoCallback(internal var context: Activity) : JavaCallback {

    override final fun sendString(data: String) {
        GlobalScope.launch(Dispatchers.Default) { // or Dispatchers.IO
            try {
                var d: CharSequence
                if (!data.contains("\n")) {
                    d = data + "\n"
                } else {
                    d = data
                }
                val i = data.length
                println("TXT: $data : $i")

                context.runOnUiThread(java.lang.Runnable {
                    context.textView.append(d)
                })

            } catch (e: Exception) {
                print("Error updating UI: ")
                println(e)
            }
        }
    }
}
