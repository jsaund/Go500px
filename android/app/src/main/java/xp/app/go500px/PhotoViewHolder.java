package xp.app.go500px;

import android.content.Context;
import android.content.Intent;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.ImageView;
import android.widget.TextView;
import com.bumptech.glide.Glide;
import go.go500px.Go500px;

public class PhotoViewHolder extends RecyclerView.ViewHolder implements View.OnClickListener {
  private final TextView mAuthor;
  private final TextView mTitle;
  private final ImageView mPhoto;
  private int mPhotoID;

  public PhotoViewHolder(View itemView) {
    super(itemView);
    itemView.setOnClickListener(this);
    mAuthor = (TextView) itemView.findViewById(R.id.photo_author);
    mTitle = (TextView) itemView.findViewById(R.id.photo_title);
    mPhoto = (ImageView) itemView.findViewById(R.id.photo);
  }

  public void bind(Go500px.Photo photo) {
    final Context context = itemView.getContext();
    mAuthor.setText(photo.User().Fullname());
    mTitle.setText(photo.Name());
    mPhotoID = photo.ID();

    Glide.with(context)
      .load(photo.Images().Get(0).URL())
      .into(mPhoto);
  }

  @Override
  public void onClick(View v) {
    final Intent photoDetailsActivity = new Intent(v.getContext(), PhotoDetailsActivity.class);
    photoDetailsActivity.putExtra(PhotoDetailsActivity.EXTRA_PHOTO_ID, String.valueOf(mPhotoID));
    v.getContext().startActivity(photoDetailsActivity);
  }
}
