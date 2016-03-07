package xp.app.go500px;

import android.os.Bundle;
import android.os.Handler;
import android.os.Message;
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

  private static class PhotoDetailsHandler extends UIHandler<PhotoDetailsActivity> {
    static final int MSG_GET_PHOTO_ERROR = 1;
    static final int MSG_GET_PHOTO_START = 2;
    static final int MSG_GET_PHOTO_SUCCESS = 3;
    static final int MSG_POST_COMMENT_ERROR = 4;
    static final int MSG_POST_COMMENT_START = 5;
    static final int MSG_POST_COMMENT_SUCCESS = 6;

    public PhotoDetailsHandler(PhotoDetailsActivity a) {
      super(a);
    }

    @Override
    void handleMessage(Message msg, PhotoDetailsActivity activity) {
      switch (msg.what) {
        case MSG_GET_PHOTO_ERROR: {
          break;
        }
        case MSG_GET_PHOTO_START: {
          final Go500px.PostCommentResponse postCommentResponse = (Go500px.PostCommentResponse) msg.obj;
          activity.mAddComment.setEnabled(true);
          Log.d(TAG, "Status: " + postCommentResponse.GetStatus() + "\nMessage: " + postCommentResponse.GetMessage() + "\nError: " + postCommentResponse.GetError());
          Toast.makeText(activity, "Posted comment!", Toast.LENGTH_LONG).show();
          break;
        }
        case MSG_GET_PHOTO_SUCCESS: {
          final Go500px.GetPhotoResponse response = (Go500px.GetPhotoResponse) msg.obj;
          final Go500px.Photo photo = response.GetPhoto();
          final Go500px.Comments comments = response.GetComments();

          final long numImages = photo.Images().Count();
          Log.d(TAG, "Number of Images: " + numImages);
          for (int i = 0; i < numImages; i++) {
            Log.d(TAG, "Image URL: " + photo.Images().Get(i).URL());
          }

          activity.mTitle.setText(photo.Name());
          activity.mAuthor.setText(photo.User().Fullname());
          activity.mDescription.setText(photo.Description());
          activity.mCamera.setText(photo.Camera());

          for (int i = 0; i < Math.min(activity.mComments.length, comments.Count()); i++) {
            final Go500px.Comment comment = comments.Get(i);
            activity.mUserComments[i].setText(comment.User().Firstname());
            activity.mComments[i].setText(comment.Body());
          }
          Glide.with(activity)
            .load(photo.Images().Get(0).URL())
            .into(activity.mPhoto);
          break;
        }
        case MSG_POST_COMMENT_ERROR: {
          activity.mAddComment.setEnabled(true);
          Toast.makeText(activity, "Failed to post comment. Reason: " + msg.obj, Toast.LENGTH_LONG).show();
          break;
        }
        case MSG_POST_COMMENT_START: {
          activity.mAddComment.setEnabled(false);
          break;
        }
        case MSG_POST_COMMENT_SUCCESS: {
          break;
        }
      }
    }
  }

  private static class GetPhotoListener extends Go500px.GetPhotoCallback.Stub {
    private final Handler mHandler;

    GetPhotoListener(Handler handler) {
      mHandler = handler;
    }

    @Override
    public void OnError(String s) {
      final Message m = mHandler.obtainMessage(PhotoDetailsHandler.MSG_GET_PHOTO_ERROR);
      m.obj = s;
      mHandler.sendMessage(m);
    }

    @Override
    public void OnStart() {
      mHandler.sendEmptyMessage(PhotoDetailsHandler.MSG_GET_PHOTO_START);
    }

    @Override
    public void OnSuccess(final Go500px.GetPhotoResponse getPhotoResponse) {
      final Message m = mHandler.obtainMessage(PhotoDetailsHandler.MSG_GET_PHOTO_SUCCESS);
      m.obj = getPhotoResponse;
      mHandler.sendMessage(m);
    }
  }

  private class PostCommentListener extends Go500px.PostCommentCallback.Stub {
    private final Handler mHandler;

    PostCommentListener(Handler handler) {
      mHandler = handler;
    }

    @Override
    public void OnError(final String s) {
      final Message m = mHandler.obtainMessage(PhotoDetailsHandler.MSG_POST_COMMENT_ERROR);
      m.obj = s;
      mHandler.sendMessage(m);
    }

    @Override
    public void OnStart() {
      mHandler.sendEmptyMessage(PhotoDetailsHandler.MSG_POST_COMMENT_START);
    }

    @Override
    public void OnSuccess(final Go500px.PostCommentResponse postCommentResponse) {
      final Message m = mHandler.obtainMessage(PhotoDetailsHandler.MSG_POST_COMMENT_SUCCESS);
      m.obj = postCommentResponse;
      mHandler.sendMessage(m);
    }
  }

  private ImageView mPhoto;
  private TextView mTitle;
  private TextView mAuthor;
  private TextView mCamera;
  private TextView mDescription;
  private TextView[] mUserComments;
  private TextView[] mComments;
  private EditText mComment;
  private Button mAddComment;

  private UIHandler mUIHandler;


  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_photo_details);
    mUIHandler = new PhotoDetailsHandler(this);

    final GetPhotoListener getPhotoListener = new GetPhotoListener(mUIHandler);
    final PostCommentListener postCommentListener = new PostCommentListener(mUIHandler);

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
        final Go500px.PostCommentRequestBuilder postCommentBuilder = Go500px.NewPostCommentRequestBuilder();
        postCommentBuilder
          .Body(comment)
          .PhotoID(photoID);
        Go500px.PostCommentAsync(postCommentBuilder, postCommentListener);
      }
    });

    final Go500px.GetPhotoRequestBuilder getPhotoBuilder = Go500px.NewGetPhotoRequestBuilder();
    final byte comments = 1;
    getPhotoBuilder
      .PhotoID(photoID)
      .Comments(comments)
      .ImageSize(4);
    Go500px.GetPhotoAsync(getPhotoBuilder, getPhotoListener);
  }

  @Override
  protected void onDestroy() {
    super.onDestroy();
    mUIHandler.removeCallbacksAndMessages(null);
  }
}
