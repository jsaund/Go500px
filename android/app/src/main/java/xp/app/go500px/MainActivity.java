package xp.app.go500px;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import go.go500px.Go500px;

public class MainActivity extends AppCompatActivity {
  private static final String TAG = "MainActivity";
  private static final String BASE_URL = "https://api.500px.com";
  private static final String CONSUMER_KEY = "8C6ImXPi4dKEnOWC3YwPnKQO1QIYbqaystDCsijC";

  private class GetPhotosListener extends Go500px.GetPhotosListener.Stub {

    @Override
    public void OnStart() {
      Log.d(TAG, "Start loading photos");
    }

    @Override
    public void OnError(String s) {
      Log.e(TAG, "Failed to retrieve photos. Reason: " + s);
    }

    @Override
    public void OnSuccess(final Go500px.Photos photos) {
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
    setContentView(R.layout.activity_main);

    mRecyclerView = (RecyclerView) findViewById(R.id.list);
    mRecyclerView.setLayoutManager(new LinearLayoutManager(this));

    Go500px.GetPhotosBuilder getPhotosBuilder = Go500px.NewGetPhotosBuilder(BASE_URL);
    getPhotosBuilder
      .Feature("popular")
      .ImageSize("20")
      .Sort("highest_rating");
    Go500px.GetPhotosAsync(getPhotosBuilder, CONSUMER_KEY, mGetPhotosListener);
  }
}
