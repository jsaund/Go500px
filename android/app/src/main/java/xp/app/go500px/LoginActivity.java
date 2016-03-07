package xp.app.go500px;

import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;
import android.os.Message;
import android.support.annotation.Nullable;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.AppCompatTextView;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import go.go500px.Go500px;
import xp.app.go500px.util.LoginHelper;

public class LoginActivity extends AppCompatActivity {
  private static final String TAG = "LoginActivity";

  private LoginCallback mLoginCallback;

  private EditText mUsername;
  private EditText mPassword;
  private AppCompatTextView mError;
  private Button mLogin;
  private ProgressDialog mProgressDialog;
  private LoginHandler mHandler;

  @Override
  protected void onCreate(@Nullable Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);

    setContentView(R.layout.activity_login);

    final FiveHundredPxCredentials credentials = LoginHelper.getCredentials(this);
    try {
      Go500px.Session(credentials.token, credentials.secret);
      launchPhotoBrowser();
    } catch (Exception e) {
      Log.d(TAG, "Authentication required");
    }

    findViewById(R.id.login_container).setVisibility(View.VISIBLE);

    mHandler = new LoginHandler(this);

    mUsername = (EditText) findViewById(R.id.email);
    mPassword = (EditText) findViewById(R.id.password);
    mError = (AppCompatTextView) findViewById(R.id.error_label);
    mLogin = (Button) findViewById(R.id.login);

    mLoginCallback = new LoginCallback();
    mLogin.setOnClickListener(new LoginClickListener());

    mProgressDialog = new ProgressDialog(LoginActivity.this, R.style.AppTheme_Dark_Dialog);
    mProgressDialog.setIndeterminate(true);
    mProgressDialog.setMessage(getString(R.string.authenticating));
  }

  private void launchPhotoBrowser() {
    final Intent launchMainActivity = new Intent(this, PhotosActivity.class);
    startActivity(launchMainActivity);
    finish();
  }

  private class LoginCallback extends Go500px.LoginCallback.Stub {

    @Override
    public void OnError(String s) {
      final Message msg = mHandler.obtainMessage(LoginHandler.MSG_LOGIN_ERROR, s);
      mHandler.sendMessage(msg);
    }

    @Override
    public void OnStart() {
      final Message msg = mHandler.obtainMessage(LoginHandler.MSG_LOGIN_START);
      mHandler.sendMessage(msg);
    }

    @Override
    public void OnSuccess(String token, String secret) {
      LoginHelper.saveLoginCredentials(LoginActivity.this, token, secret);
      final Message msg = mHandler.obtainMessage(LoginHandler.MSG_LOGIN_SUCCESS);
      mHandler.sendMessage(msg);
    }
  }

  private class LoginClickListener implements View.OnClickListener {

    @Override
    public void onClick(View v) {
      final String username = mUsername.getText().toString();
      final String password = mPassword.getText().toString();
      Go500px.Login(username, password, mLoginCallback);
    }
  }

  private static class LoginHandler extends UIHandler<LoginActivity> {
    private static final int MSG_LOGIN_START = 1;
    private static final int MSG_LOGIN_ERROR = 2;
    private static final int MSG_LOGIN_SUCCESS = 3;

    LoginHandler(LoginActivity activity) {
      super(activity);
    }

    @Override
    public void handleMessage(Message msg, LoginActivity activity) {
      switch (msg.what) {
        case MSG_LOGIN_START:
          activity.mError.setVisibility(View.GONE);
          activity.mLogin.setEnabled(false);
          activity.mProgressDialog.show();
          break;
        case MSG_LOGIN_ERROR:
          activity.mError.setVisibility(View.VISIBLE);
          activity.mError.setText(msg.obj.toString());
          activity.mLogin.setEnabled(true);
          activity.mProgressDialog.cancel();
          break;
        case MSG_LOGIN_SUCCESS:
          activity.mError.setVisibility(View.GONE);
          activity.mProgressDialog.cancel();
          activity.launchPhotoBrowser();
          break;
      }
    }
  }
}
