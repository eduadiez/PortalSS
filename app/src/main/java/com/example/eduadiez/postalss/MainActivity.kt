package com.example.eduadiez.postalss


import android.content.Context
import android.widget.Toast
import android.widget.*
import android.support.v7.app.AppCompatActivity
import android.os.Bundle
import android.support.design.widget.TextInputEditText
import gomobile.Gomobile
import kotlinx.android.synthetic.main.activity_main.*
import android.text.method.ScrollingMovementMethod
import android.support.design.widget.TextInputLayout
import android.util.Log
import kotlinx.coroutines.*
import java.security.AccessController.getContext
import kotlin.coroutines.CoroutineContext


class MainActivity : AppCompatActivity(), CoroutineScope {

    override val coroutineContext: CoroutineContext
        get() = Dispatchers.Main

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        // scrollview
        val tview: TextView = findViewById(R.id.textView)
        val smm = ScrollingMovementMethod()
        tview.movementMethod = smm

        // callback
        val gocb = GoCallback(this)
        Gomobile.registerJavaCallback(gocb)

        startButton.setOnClickListener{
            GlobalScope.launch(Dispatchers.Main) {
                toast("Launching SWARM PPS process!")
                Log.i("PATH",getApplicationInfo().dataDir)
                val tviewnick: TextView = findViewById(R.id.editTextNick)
                val tviewtopic: TextView = findViewById(R.id.editTextTopic)
                withContext(Dispatchers.Default) { Gomobile.starting(getApplicationInfo().dataDir,tviewtopic.text.toString(),tviewnick.text.toString()) }
                startButton.isEnabled = false
                tviewnick.isEnabled = false
                tviewtopic.isEnabled = false

            }
        }

        sendButton.setOnClickListener {
            GlobalScope.launch(Dispatchers.Main) {
                val tiet = findViewById(R.id.textInputEditText) as TextInputEditText
                GlobalScope.launch(Dispatchers.Default) { Gomobile.sendMessage( tiet.text.toString()) }
                tiet.setText("");

            }
        }
    }

    // creates toast messages to be displayed
    public fun Context.toast(message: CharSequence) =
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show()


}