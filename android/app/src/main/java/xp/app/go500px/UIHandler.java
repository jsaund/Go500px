package xp.app.go500px;

import android.app.Activity;
import android.os.Handler;
import android.os.Message;

import java.lang.ref.WeakReference;

public abstract class UIHandler<T extends Activity> extends Handler {
  private final WeakReference<T> mRef;

  public UIHandler(T a) {
    mRef = new WeakReference<>(a);
  }

  @Override
  public void handleMessage(Message msg) {
    final T a = mRef.get();
    if (a == null || a.isFinishing()) {
      return;
    }
    handleMessage(msg, a);
  }

  abstract void handleMessage(Message msg, T activity);
}
