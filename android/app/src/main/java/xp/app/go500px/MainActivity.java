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

  @Override
  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);

    RecyclerView recyclerView = (RecyclerView) findViewById(R.id.list);
    recyclerView.setLayoutManager(new LinearLayoutManager(this));

    try {
      Go500px.GetPhotosBuilder getPhotosBuilder = Go500px.NewGetPhotosBuilder(BASE_URL);
      getPhotosBuilder
        .Feature("popular")
        .ImageSize("20")
        .Sort("highest_rating");
      Go500px.Photos photos = Go500px.GetPhotos(getPhotosBuilder,CONSUMER_KEY);
      final int numPhotos = photos.Count();
      Log.d(TAG, "Received " + numPhotos + " photos!");

      recyclerView.setAdapter(new PhotosAdapter(photos));
    } catch (Exception e) {
      Log.e(TAG, "Failed to retrieve photos.", e);
    }
  }
}
