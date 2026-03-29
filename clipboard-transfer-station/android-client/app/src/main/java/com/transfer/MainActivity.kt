package com.transfer

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.runtime.Composable
import androidx.compose.ui.tooling.preview.Preview
import com.transfer.ui.theme.ClipboardTransferStationTheme

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            ClipboardTransferStationTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    color = MaterialTheme.colorScheme.background
                ) {
                    Greeting("Welcome to Clipboard Transfer Station")
                }
            }
        }
    }
}

@Composable
fun Greeting(name: String) {
    // TODO: Implement the UI for the main activity
}

@Preview(showBackground = true)
@Composable
fun DefaultPreview() {
    ClipboardTransferStationTheme {
        Greeting("Welcome to Clipboard Transfer Station")
    }
}