package xp.app.go500px;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import go.go500px.Go500px;

public class PhotosActivity extends AppCompatActivity {
  private static final String TAG = "PhotosActivity";

  private class GetPhotosListener extends Go500px.GetPhotosCallback.Stub {

    @Override
    public void OnError(String s) {
      Log.e(TAG, "Failed to retrieve photos. Reason: " + s);
    }

    @Override
    public void OnStart() {
      Log.d(TAG, "Start loading photos");
    }

    @Override
    public void OnSuccess(Go500px.GetPhotosResponse getPhotosResponse) {
      final Go500px.Photos photos = getPhotosResponse.GetPhotos();
      final int numPhotos = photos.Count();
      Log.d(TAG, "Received " + numPhotos + " photos!");

      runOnUiThread(new Runnable() {
        public void run() {
          mRecyclerView.setAdapter(new PhotosAdapter(photos));
        }
      });
    }
  }

  private final GetPhotosListener mGetPhotosListener = new GetPhotosListener();
  private RecyclerView mRecyclerView;

  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_photos);

    mRecyclerView = (RecyclerView) findViewById(R.id.list);
    mRecyclerView.setLayoutManager(new LinearLayoutManager(this));

    final Go500px.GetPhotosRequestBuilder getPhotosBuilder = Go500px.NewGetPhotosRequestBuilder(Config.BASE_URL);
    getPhotosBuilder
      .Feature("popular")
      .ImageSize("20")
      .Sort("highest_rating");
    Go500px.GetPhotosAsync(getPhotosBuilder, mGetPhotosListener);
  }
}
