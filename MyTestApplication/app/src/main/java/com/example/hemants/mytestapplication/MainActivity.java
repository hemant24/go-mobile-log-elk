package com.example.hemants.mytestapplication;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import rlogger.Rlogger;
import rlogger.Logstash;


public class MainActivity extends AppCompatActivity implements View.OnClickListener  {

    String tcpHost = "0.tcp.ngrok.io";
    int tcpPort = 13632;

    Logstash logger;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.logger);
        Button sendLogBtn = (Button) findViewById(R.id.button);
        sendLogBtn.setOnClickListener(this);
    }

    @Override
    public void onClick(View view) {
        String logMessage = ((EditText)findViewById(R.id.editText)).getText().toString();
        logger = Rlogger.new_(tcpHost, tcpPort , 20);
        logger.connect();
        if(logger != null) {
            boolean writeStatus = logger.writeln("android", "INFO", "test-tag", logMessage);
            Toast.makeText(this, "Status of write : " + writeStatus, Toast.LENGTH_SHORT);
        }else {
            Toast.makeText(this, "Logger failed to initialise", Toast.LENGTH_SHORT);
        }

    }
}
