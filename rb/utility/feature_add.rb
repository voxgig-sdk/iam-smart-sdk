# IamSmart SDK utility: feature_add
module IamSmartUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
