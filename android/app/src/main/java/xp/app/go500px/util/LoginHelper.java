package xp.app.go500px.util;

import android.content.Context;
import android.content.SharedPreferences;
import xp.app.go500px.FiveHundredPxCredentials;
import xp.app.go500px.R;

public final class LoginHelper {

  public static void saveLoginCredentials(Context context, String token, String secret) {
    final SharedPreferences sharedPreferences = context.getSharedPreferences(context.getString(R.string.prefs_login), Context.MODE_PRIVATE);
    final SharedPreferences.Editor editor = sharedPreferences.edit();
    editor.putString(context.getString(R.string.key_token), token);
    editor.putString(context.getString(R.string.key_secret), secret);
    editor.commit();
  }

  public static FiveHundredPxCredentials getCredentials(Context context) {
    final SharedPreferences sharedPreferences = context.getSharedPreferences(context.getString(R.string.prefs_login), Context.MODE_PRIVATE);
    final String token = sharedPreferences.getString(context.getString(R.string.key_token), null);
    final String secret = sharedPreferences.getString(context.getString(R.string.key_secret), null);
    return new FiveHundredPxCredentials(token, secret);
  }
}
