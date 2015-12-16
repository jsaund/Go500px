package xp.app.go500px;

import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import go.go500px.Go500px;

public class PhotosAdapter extends RecyclerView.Adapter<PhotoViewHolder> {
  private Go500px.Photos mPhotos;

  public PhotosAdapter(Go500px.Photos photos) {
    mPhotos = photos;
  }

  @Override
  public PhotoViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
    final View itemView = LayoutInflater.from(parent.getContext()).inflate(R.layout.photo_row_item, parent, false);
    return new PhotoViewHolder(itemView);
  }

  @Override
  public void onBindViewHolder(PhotoViewHolder holder, int position) {
    holder.bind(mPhotos.Get(position));
  }

  @Override
  public int getItemCount() {
    return mPhotos.Count();
  }
}
