package xp.app.go500px;

import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import go.go500px.Go500px;

public class PhotosActivity extends AppCompatActivity {
  private static final String TAG = "PhotosActivity";

  private static class GetPhotosListener extends Go500px.GetPhotosCallback.Stub {
    private final Handler mHandler;

    GetPhotosListener(Handler handler) {
      mHandler = handler;
    }

    @Override
    public void OnError(String s) {
      final Message m = mHandler.obtainMessage(PhotosHandler.MSG_ERROR);
      m.obj = s;
      mHandler.sendMessage(m);
    }

    @Override
    public void OnStart() {
      mHandler.sendEmptyMessage(PhotosHandler.MSG_START);
    }

    @Override
    public void OnSuccess(Go500px.GetPhotosResponse getPhotosResponse) {
      final Go500px.Photos photos = getPhotosResponse.GetPhotos();
      final int numPhotos = photos.Count();
      Log.d(TAG, "Received " + numPhotos + " photos!");

      final Message m = mHandler.obtainMessage(PhotosHandler.MSG_SUCCESS);
      m.obj = photos;
      mHandler.sendMessage(m);
    }
  }

  private static class PhotosHandler extends UIHandler<PhotosActivity> {
    static final int MSG_ERROR = 1;
    static final int MSG_START = 2;
    static final int MSG_SUCCESS = 3;

    PhotosHandler(PhotosActivity a) {
      super(a);
    }

    @Override
    public void handleMessage(Message msg, PhotosActivity activity) {
      switch (msg.what) {
        case MSG_ERROR: {
          Log.e(TAG, "Error fetching photos: " + msg.obj);
          break;
        }
        case MSG_START: {
          Log.d(TAG, "Fetching photos");
          break;
        }
        case MSG_SUCCESS: {
          final Go500px.Photos photos = (Go500px.Photos) msg.obj;
          activity.mRecyclerView.setAdapter(new PhotosAdapter(photos));
          break;
        }
      }
    }
  }

  private RecyclerView mRecyclerView;
  private Handler mHandler;

  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_photos);

    mRecyclerView = (RecyclerView) findViewById(R.id.list);
    mRecyclerView.setLayoutManager(new LinearLayoutManager(this));

    mHandler = new PhotosHandler(this);
    final GetPhotosListener getPhotosListener = new GetPhotosListener(mHandler);
    final Go500px.GetPhotosRequestBuilder getPhotosBuilder = Go500px.NewGetPhotosRequestBuilder();
    getPhotosBuilder
      .Feature("popular")
      .ImageSize("20")
      .Sort("highest_rating");
    Go500px.GetPhotosAsync(getPhotosBuilder, getPhotosListener);
  }

  @Override
  protected void onDestroy() {
    super.onDestroy();
    mHandler.removeCallbacksAndMessages(null);
  }
}
