package xp.app.go500px;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.text.TextUtils;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;
import com.bumptech.glide.Glide;
import go.go500px.Go500px;

public class PhotoDetailsActivity extends AppCompatActivity {
  private static final String TAG = "PhotoDetailsActivity";

  public static final String EXTRA_PHOTO_ID = "xp.app.go500px.PhotoDetailsActivity.EXTRA_PHOTO_ID";

  private class GetPhotoListener extends Go500px.GetPhotoCallback.Stub {

    @Override
    public void OnError(String s) {
    }

    @Override
    public void OnStart() {
    }

    @Override
    public void OnSuccess(final Go500px.GetPhotoResponse getPhotoResponse) {
      final Go500px.Photo photo = getPhotoResponse.GetPhoto();
      final Go500px.Comments comments = getPhotoResponse.GetComments();
      final long numImages = photo.Images().Count();
      Log.d(TAG, "Number of Images: " + numImages);
      for (int i = 0; i < numImages; i++) {
        Log.d(TAG, "Image URL: " + photo.Images().Get(i).URL());
      }

      runOnUiThread(new Runnable() {
        @Override
        public void run() {
          mTitle.setText(photo.Name());
          mAuthor.setText(photo.User().Fullname());
          mDescription.setText(photo.Description());
          mCamera.setText(photo.Camera());

          for (int i = 0; i < Math.min(mComments.length, comments.Count()); i++) {
            final Go500px.Comment comment = comments.Get(i);
            mUserComments[i].setText(comment.User().Firstname());
            mComments[i].setText(comment.Body());
          }
          Glide.with(PhotoDetailsActivity.this)
            .load(photo.Images().Get(0).URL())
            .into(mPhoto);
        }
      });
    }
  }

  private class PostCommentListener extends Go500px.PostCommentCallback.Stub {

    @Override
    public void OnError(final String s) {
      runOnUiThread(new Runnable() {
        @Override
        public void run() {
          mAddComment.setEnabled(true);
          Toast.makeText(PhotoDetailsActivity.this, "Failed to post comment. Reason: " + s, Toast.LENGTH_LONG).show();
        }
      });
    }

    @Override
    public void OnStart() {
      runOnUiThread(new Runnable() {
        @Override
        public void run() {
          mAddComment.setEnabled(false);
        }
      });
    }

    @Override
    public void OnSuccess(final Go500px.PostCommentResponse postCommentResponse) {
      runOnUiThread(new Runnable() {
        @Override
        public void run() {
          mAddComment.setEnabled(true);
          Log.d(TAG, "Status: " + postCommentResponse.GetStatus() + "\nMessage: " + postCommentResponse.GetMessage() + "\nError: " + postCommentResponse.GetError());
          Toast.makeText(PhotoDetailsActivity.this, "Posted comment!", Toast.LENGTH_LONG).show();
        }
      });
    }
  }

  private final GetPhotoListener mGetPhotoListener = new GetPhotoListener();
  private final PostCommentListener mPostCommentListener = new PostCommentListener();

  private ImageView mPhoto;
  private TextView mTitle;
  private TextView mAuthor;
  private TextView mCamera;
  private TextView mDescription;
  private TextView[] mUserComments;
  private TextView[] mComments;
  private EditText mComment;
  private Button mAddComment;


  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_photo_details);

    final String photoID = getIntent().getStringExtra(EXTRA_PHOTO_ID);

    mPhoto = (ImageView) findViewById(R.id.photo);
    mTitle = (TextView) findViewById(R.id.photo_title);
    mAuthor= (TextView) findViewById(R.id.photo_author);
    mCamera= (TextView) findViewById(R.id.photo_camera);
    mDescription= (TextView) findViewById(R.id.photo_description);

    mUserComments = new TextView[2];
    mUserComments[0] = (TextView) findViewById(R.id.photo_user_comments_1);
    mUserComments[1] = (TextView) findViewById(R.id.photo_user_comments_2);

    mComments = new TextView[2];
    mComments[0]= (TextView) findViewById(R.id.photo_comments_1);
    mComments[1]= (TextView) findViewById(R.id.photo_comments_2);

    mComment = (EditText) findViewById(R.id.comment);
    mAddComment = (Button) findViewById(R.id.add_comment);
    mAddComment.setOnClickListener(new View.OnClickListener() {
      @Override
      public void onClick(View v) {
        final String comment = mComment.getText().toString();
        if (TextUtils.isEmpty(comment)) {
          return;
        }
        final Go500px.PostCommentRequestBuilder postCommentBuilder = Go500px.NewPostCommentRequestBuilder(Config.BASE_URL);
        postCommentBuilder
          .Body(comment)
          .PhotoID(photoID);
        Go500px.PostCommentAsync(postCommentBuilder, mPostCommentListener);
      }
    });

    final Go500px.GetPhotoRequestBuilder getPhotoBuilder = Go500px.NewGetPhotoRequestBuilder(Config.BASE_URL);
    final byte comments = 1;
    getPhotoBuilder
      .PhotoID(photoID)
      .Comments(comments)
      .ImageSize(4);
    Go500px.GetPhotoAsync(getPhotoBuilder, mGetPhotoListener);
  }

}
