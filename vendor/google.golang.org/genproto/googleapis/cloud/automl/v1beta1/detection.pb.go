// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/detection.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Annotation details for image object detection.
type ImageObjectDetectionAnnotation struct {
	// Output only.
	// The rectangle representing the object location.
	BoundingBox *BoundingPoly `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Output only.
	// The confidence that this annotation is positive for the parent example,
	// value in [0, 1], higher means higher positivity confidence.
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageObjectDetectionAnnotation) Reset()         { *m = ImageObjectDetectionAnnotation{} }
func (m *ImageObjectDetectionAnnotation) String() string { return proto.CompactTextString(m) }
func (*ImageObjectDetectionAnnotation) ProtoMessage()    {}
func (*ImageObjectDetectionAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{0}
}

func (m *ImageObjectDetectionAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageObjectDetectionAnnotation.Unmarshal(m, b)
}
func (m *ImageObjectDetectionAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageObjectDetectionAnnotation.Marshal(b, m, deterministic)
}
func (m *ImageObjectDetectionAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageObjectDetectionAnnotation.Merge(m, src)
}
func (m *ImageObjectDetectionAnnotation) XXX_Size() int {
	return xxx_messageInfo_ImageObjectDetectionAnnotation.Size(m)
}
func (m *ImageObjectDetectionAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageObjectDetectionAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ImageObjectDetectionAnnotation proto.InternalMessageInfo

func (m *ImageObjectDetectionAnnotation) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *ImageObjectDetectionAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Annotation details for video object tracking.
type VideoObjectTrackingAnnotation struct {
	// Optional.
	// The instance of the object, expressed as a positive integer. Used to tell
	// apart objects of the same type (i.e. AnnotationSpec) when multiple are
	// present on a single example.
	// NOTE: Instance ID prediction quality is not a part of model evaluation and
	// is done as best effort. Especially in cases when an entity goes
	// off-screen for a longer time (minutes), when it comes back it may be given
	// a new instance ID.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Required. A time (frame) of a video to which this annotation pertains.
	// Represented as the duration since the video's start.
	TimeOffset *duration.Duration `protobuf:"bytes,2,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	// Required. The rectangle representing the object location on the frame (i.e.
	// at the time_offset of the video).
	BoundingBox *BoundingPoly `protobuf:"bytes,3,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	// Output only.
	// The confidence that this annotation is positive for the video at
	// the time_offset, value in [0, 1], higher means higher positivity
	// confidence. For annotations created by the user the score is 1. When
	// user approves an annotation, the original float score is kept (and not
	// changed to 1).
	Score                float32  `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoObjectTrackingAnnotation) Reset()         { *m = VideoObjectTrackingAnnotation{} }
func (m *VideoObjectTrackingAnnotation) String() string { return proto.CompactTextString(m) }
func (*VideoObjectTrackingAnnotation) ProtoMessage()    {}
func (*VideoObjectTrackingAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{1}
}

func (m *VideoObjectTrackingAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoObjectTrackingAnnotation.Unmarshal(m, b)
}
func (m *VideoObjectTrackingAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoObjectTrackingAnnotation.Marshal(b, m, deterministic)
}
func (m *VideoObjectTrackingAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoObjectTrackingAnnotation.Merge(m, src)
}
func (m *VideoObjectTrackingAnnotation) XXX_Size() int {
	return xxx_messageInfo_VideoObjectTrackingAnnotation.Size(m)
}
func (m *VideoObjectTrackingAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoObjectTrackingAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_VideoObjectTrackingAnnotation proto.InternalMessageInfo

func (m *VideoObjectTrackingAnnotation) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *VideoObjectTrackingAnnotation) GetTimeOffset() *duration.Duration {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func (m *VideoObjectTrackingAnnotation) GetBoundingBox() *BoundingPoly {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func (m *VideoObjectTrackingAnnotation) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Bounding box matching model metrics for a single intersection-over-union
// threshold and multiple label match confidence thresholds.
type BoundingBoxMetricsEntry struct {
	// Output only. The intersection-over-union threshold value used to compute
	// this metrics entry.
	IouThreshold float32 `protobuf:"fixed32,1,opt,name=iou_threshold,json=iouThreshold,proto3" json:"iou_threshold,omitempty"`
	// Output only. The mean average precision, most often close to au_prc.
	MeanAveragePrecision float32 `protobuf:"fixed32,2,opt,name=mean_average_precision,json=meanAveragePrecision,proto3" json:"mean_average_precision,omitempty"`
	// Output only. Metrics for each label-match confidence_threshold from
	// 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99. Precision-recall curve is
	// derived from them.
	ConfidenceMetricsEntries []*BoundingBoxMetricsEntry_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entries,json=confidenceMetricsEntries,proto3" json:"confidence_metrics_entries,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                          `json:"-"`
	XXX_unrecognized         []byte                                            `json:"-"`
	XXX_sizecache            int32                                             `json:"-"`
}

func (m *BoundingBoxMetricsEntry) Reset()         { *m = BoundingBoxMetricsEntry{} }
func (m *BoundingBoxMetricsEntry) String() string { return proto.CompactTextString(m) }
func (*BoundingBoxMetricsEntry) ProtoMessage()    {}
func (*BoundingBoxMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{2}
}

func (m *BoundingBoxMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingBoxMetricsEntry.Unmarshal(m, b)
}
func (m *BoundingBoxMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingBoxMetricsEntry.Marshal(b, m, deterministic)
}
func (m *BoundingBoxMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBoxMetricsEntry.Merge(m, src)
}
func (m *BoundingBoxMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_BoundingBoxMetricsEntry.Size(m)
}
func (m *BoundingBoxMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBoxMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBoxMetricsEntry proto.InternalMessageInfo

func (m *BoundingBoxMetricsEntry) GetIouThreshold() float32 {
	if m != nil {
		return m.IouThreshold
	}
	return 0
}

func (m *BoundingBoxMetricsEntry) GetMeanAveragePrecision() float32 {
	if m != nil {
		return m.MeanAveragePrecision
	}
	return 0
}

func (m *BoundingBoxMetricsEntry) GetConfidenceMetricsEntries() []*BoundingBoxMetricsEntry_ConfidenceMetricsEntry {
	if m != nil {
		return m.ConfidenceMetricsEntries
	}
	return nil
}

// Metrics for a single confidence threshold.
type BoundingBoxMetricsEntry_ConfidenceMetricsEntry struct {
	// Output only. The confidence threshold value used to compute the metrics.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Recall under the given confidence threshold.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision under the given confidence threshold.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score              float32  `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) Reset() {
	*m = BoundingBoxMetricsEntry_ConfidenceMetricsEntry{}
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) String() string {
	return proto.CompactTextString(m)
}
func (*BoundingBoxMetricsEntry_ConfidenceMetricsEntry) ProtoMessage() {}
func (*BoundingBoxMetricsEntry_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{2, 0}
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Unmarshal(m, b)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Marshal(b, m, deterministic)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Merge(m, src)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.Size(m)
}
func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBoxMetricsEntry_ConfidenceMetricsEntry proto.InternalMessageInfo

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if m != nil {
		return m.ConfidenceThreshold
	}
	return 0
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *BoundingBoxMetricsEntry_ConfidenceMetricsEntry) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

// Model evaluation metrics for image object detection problems.
// Evaluates prediction quality of labeled bounding boxes.
type ImageObjectDetectionEvaluationMetrics struct {
	// Output only. The total number of bounding boxes (i.e. summed over all
	// images) the ground truth used to create this evaluation had.
	EvaluatedBoundingBoxCount int32 `protobuf:"varint,1,opt,name=evaluated_bounding_box_count,json=evaluatedBoundingBoxCount,proto3" json:"evaluated_bounding_box_count,omitempty"`
	// Output only. The bounding boxes match metrics for each
	// Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// pair.
	BoundingBoxMetricsEntries []*BoundingBoxMetricsEntry `protobuf:"bytes,2,rep,name=bounding_box_metrics_entries,json=boundingBoxMetricsEntries,proto3" json:"bounding_box_metrics_entries,omitempty"`
	// Output only. The single metric for bounding boxes evaluation:
	// the mean_average_precision averaged over all bounding_box_metrics_entries.
	BoundingBoxMeanAveragePrecision float32  `protobuf:"fixed32,3,opt,name=bounding_box_mean_average_precision,json=boundingBoxMeanAveragePrecision,proto3" json:"bounding_box_mean_average_precision,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *ImageObjectDetectionEvaluationMetrics) Reset()         { *m = ImageObjectDetectionEvaluationMetrics{} }
func (m *ImageObjectDetectionEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*ImageObjectDetectionEvaluationMetrics) ProtoMessage()    {}
func (*ImageObjectDetectionEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{3}
}

func (m *ImageObjectDetectionEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Unmarshal(m, b)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Merge(m, src)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.Size(m)
}
func (m *ImageObjectDetectionEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageObjectDetectionEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ImageObjectDetectionEvaluationMetrics proto.InternalMessageInfo

func (m *ImageObjectDetectionEvaluationMetrics) GetEvaluatedBoundingBoxCount() int32 {
	if m != nil {
		return m.EvaluatedBoundingBoxCount
	}
	return 0
}

func (m *ImageObjectDetectionEvaluationMetrics) GetBoundingBoxMetricsEntries() []*BoundingBoxMetricsEntry {
	if m != nil {
		return m.BoundingBoxMetricsEntries
	}
	return nil
}

func (m *ImageObjectDetectionEvaluationMetrics) GetBoundingBoxMeanAveragePrecision() float32 {
	if m != nil {
		return m.BoundingBoxMeanAveragePrecision
	}
	return 0
}

// Model evaluation metrics for video object tracking problems.
// Evaluates prediction quality of both labeled bounding boxes and labeled
// tracks (i.e. series of bounding boxes sharing same label and instance ID).
type VideoObjectTrackingEvaluationMetrics struct {
	// Output only. The number of video frames used to create this evaluation.
	EvaluatedFrameCount int32 `protobuf:"varint,1,opt,name=evaluated_frame_count,json=evaluatedFrameCount,proto3" json:"evaluated_frame_count,omitempty"`
	// Output only. The total number of bounding boxes (i.e. summed over all
	// frames) the ground truth used to create this evaluation had.
	EvaluatedBoundingBoxCount int32 `protobuf:"varint,2,opt,name=evaluated_bounding_box_count,json=evaluatedBoundingBoxCount,proto3" json:"evaluated_bounding_box_count,omitempty"`
	// Output only. The bounding boxes match metrics for each
	// Intersection-over-union threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// and each label confidence threshold 0.05,0.10,...,0.95,0.96,0.97,0.98,0.99
	// pair.
	BoundingBoxMetricsEntries []*BoundingBoxMetricsEntry `protobuf:"bytes,4,rep,name=bounding_box_metrics_entries,json=boundingBoxMetricsEntries,proto3" json:"bounding_box_metrics_entries,omitempty"`
	// Output only. The single metric for bounding boxes evaluation:
	// the mean_average_precision averaged over all bounding_box_metrics_entries.
	BoundingBoxMeanAveragePrecision float32  `protobuf:"fixed32,6,opt,name=bounding_box_mean_average_precision,json=boundingBoxMeanAveragePrecision,proto3" json:"bounding_box_mean_average_precision,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *VideoObjectTrackingEvaluationMetrics) Reset()         { *m = VideoObjectTrackingEvaluationMetrics{} }
func (m *VideoObjectTrackingEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*VideoObjectTrackingEvaluationMetrics) ProtoMessage()    {}
func (*VideoObjectTrackingEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_85c92c72aeb2dfbd, []int{4}
}

func (m *VideoObjectTrackingEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Unmarshal(m, b)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Merge(m, src)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.Size(m)
}
func (m *VideoObjectTrackingEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoObjectTrackingEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_VideoObjectTrackingEvaluationMetrics proto.InternalMessageInfo

func (m *VideoObjectTrackingEvaluationMetrics) GetEvaluatedFrameCount() int32 {
	if m != nil {
		return m.EvaluatedFrameCount
	}
	return 0
}

func (m *VideoObjectTrackingEvaluationMetrics) GetEvaluatedBoundingBoxCount() int32 {
	if m != nil {
		return m.EvaluatedBoundingBoxCount
	}
	return 0
}

func (m *VideoObjectTrackingEvaluationMetrics) GetBoundingBoxMetricsEntries() []*BoundingBoxMetricsEntry {
	if m != nil {
		return m.BoundingBoxMetricsEntries
	}
	return nil
}

func (m *VideoObjectTrackingEvaluationMetrics) GetBoundingBoxMeanAveragePrecision() float32 {
	if m != nil {
		return m.BoundingBoxMeanAveragePrecision
	}
	return 0
}

func init() {
	proto.RegisterType((*ImageObjectDetectionAnnotation)(nil), "google.cloud.automl.v1beta1.ImageObjectDetectionAnnotation")
	proto.RegisterType((*VideoObjectTrackingAnnotation)(nil), "google.cloud.automl.v1beta1.VideoObjectTrackingAnnotation")
	proto.RegisterType((*BoundingBoxMetricsEntry)(nil), "google.cloud.automl.v1beta1.BoundingBoxMetricsEntry")
	proto.RegisterType((*BoundingBoxMetricsEntry_ConfidenceMetricsEntry)(nil), "google.cloud.automl.v1beta1.BoundingBoxMetricsEntry.ConfidenceMetricsEntry")
	proto.RegisterType((*ImageObjectDetectionEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.ImageObjectDetectionEvaluationMetrics")
	proto.RegisterType((*VideoObjectTrackingEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.VideoObjectTrackingEvaluationMetrics")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/detection.proto", fileDescriptor_85c92c72aeb2dfbd)
}

var fileDescriptor_85c92c72aeb2dfbd = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0xd1, 0x4e, 0xdb, 0x3a,
	0x18, 0xc7, 0x95, 0x14, 0x38, 0x07, 0x97, 0x73, 0x13, 0x38, 0x9c, 0xb6, 0xf4, 0x00, 0x82, 0x4d,
	0x62, 0x9b, 0x94, 0xa8, 0x1d, 0x57, 0xdb, 0xc5, 0xd4, 0x02, 0x9b, 0xd0, 0x40, 0xa0, 0x0c, 0x71,
	0x31, 0x21, 0x45, 0x8e, 0xf3, 0x25, 0x78, 0x4b, 0xec, 0xca, 0x71, 0x10, 0xdc, 0x73, 0xb3, 0x77,
	0x98, 0xb4, 0x27, 0xd8, 0x8b, 0xec, 0x29, 0x78, 0x94, 0xa9, 0x76, 0xd2, 0x84, 0x12, 0x31, 0xb4,
	0x4d, 0xda, 0xa5, 0xbf, 0xff, 0xdf, 0x9f, 0xff, 0xdf, 0xcf, 0x4e, 0x8b, 0x9e, 0x45, 0x9c, 0x47,
	0x31, 0x38, 0x24, 0xe6, 0x59, 0xe0, 0xe0, 0x4c, 0xf2, 0x24, 0x76, 0x2e, 0x7a, 0x3e, 0x48, 0xdc,
	0x73, 0x02, 0x90, 0x40, 0x24, 0xe5, 0xcc, 0x1e, 0x09, 0x2e, 0xb9, 0xb5, 0xa2, 0xcd, 0xb6, 0x32,
	0xdb, 0xda, 0x6c, 0xe7, 0xe6, 0xce, 0xd3, 0xfb, 0x3a, 0x45, 0xc0, 0x13, 0x90, 0xe2, 0x4a, 0x37,
	0xea, 0xac, 0xe6, 0x5e, 0xb5, 0xf2, 0xb3, 0xd0, 0x09, 0x32, 0x81, 0xcb, 0x83, 0x3a, 0xdd, 0x5c,
	0xc7, 0x23, 0xea, 0x60, 0xc6, 0xb8, 0x54, 0x62, 0xaa, 0xd5, 0x8d, 0x6b, 0x03, 0xad, 0xee, 0x27,
	0x38, 0x82, 0x23, 0xff, 0x03, 0x10, 0xb9, 0x5b, 0xa4, 0x1c, 0x4c, 0x9c, 0xd6, 0x01, 0x5a, 0xf0,
	0x79, 0xc6, 0x02, 0xca, 0x22, 0xcf, 0xe7, 0x97, 0x2d, 0x63, 0xdd, 0xd8, 0x6a, 0xf6, 0x9f, 0xd8,
	0xf7, 0x0c, 0x60, 0x0f, 0xf3, 0x0d, 0xc7, 0x3c, 0xbe, 0x72, 0x9b, 0xc5, 0xf6, 0x21, 0xbf, 0xb4,
	0x96, 0xd0, 0x6c, 0x4a, 0xb8, 0x80, 0x96, 0xb9, 0x6e, 0x6c, 0x99, 0xae, 0x5e, 0x6c, 0xdc, 0x18,
	0xe8, 0xff, 0x53, 0x1a, 0x00, 0xd7, 0x31, 0x4e, 0x04, 0x26, 0x1f, 0x29, 0x8b, 0x2a, 0x29, 0xd6,
	0x50, 0x93, 0xb2, 0x54, 0x62, 0x46, 0xc0, 0xa3, 0x81, 0x0a, 0x31, 0xef, 0xa2, 0xa2, 0xb4, 0x1f,
	0x58, 0x2f, 0x50, 0x53, 0xd2, 0x04, 0x3c, 0x1e, 0x86, 0x29, 0x48, 0xd5, 0xbe, 0xd9, 0x6f, 0x17,
	0x29, 0x0b, 0x3a, 0xf6, 0x6e, 0x4e, 0xc7, 0x45, 0x63, 0xf7, 0x91, 0x32, 0xdf, 0x19, 0xb1, 0xf1,
	0x7b, 0x46, 0x9c, 0xa9, 0x8e, 0xf8, 0xa5, 0x81, 0xfe, 0x1b, 0x96, 0xae, 0x43, 0x90, 0x82, 0x92,
	0x74, 0x8f, 0x49, 0x71, 0x65, 0x6d, 0xa2, 0x7f, 0x28, 0xcf, 0x3c, 0x79, 0x2e, 0x20, 0x3d, 0xe7,
	0xb1, 0x1e, 0xcf, 0x74, 0x17, 0x28, 0xcf, 0x4e, 0x8a, 0x9a, 0xb5, 0x8d, 0x96, 0x13, 0xc0, 0xcc,
	0xc3, 0x17, 0x20, 0x70, 0x04, 0xde, 0x48, 0x00, 0xa1, 0x29, 0xe5, 0x2c, 0x47, 0xb9, 0x34, 0x56,
	0x07, 0x5a, 0x3c, 0x2e, 0x34, 0xeb, 0x93, 0x81, 0x3a, 0x84, 0xb3, 0x90, 0x06, 0x30, 0x46, 0x97,
	0xe8, 0x63, 0x3d, 0x60, 0x52, 0x50, 0x48, 0x5b, 0x8d, 0xf5, 0xc6, 0x56, 0xb3, 0xff, 0xf6, 0x41,
	0x93, 0x4e, 0xa5, 0xb6, 0x77, 0x26, 0x6d, 0xab, 0x65, 0xb7, 0x45, 0xea, 0xea, 0x14, 0xd2, 0xce,
	0x67, 0x03, 0x2d, 0xd7, 0x6f, 0xb2, 0x7a, 0x68, 0xa9, 0x92, 0x72, 0x1a, 0xc4, 0x62, 0xa9, 0x95,
	0x3c, 0x96, 0xd1, 0x9c, 0x00, 0x82, 0xe3, 0x38, 0x9f, 0x3f, 0x5f, 0x59, 0x5d, 0x34, 0x5f, 0xa2,
	0x69, 0x28, 0xa9, 0x2c, 0x58, 0x6d, 0xf4, 0x77, 0xd8, 0xf3, 0xaa, 0xf7, 0xf3, 0x57, 0xd8, 0x7b,
	0xa7, 0x6e, 0xe8, 0xab, 0x89, 0x1e, 0xd7, 0x7d, 0x0b, 0x7b, 0x17, 0x38, 0xce, 0xd4, 0xa3, 0xc9,
	0x23, 0x5b, 0xaf, 0x50, 0x17, 0x74, 0x11, 0x02, 0xaf, 0xfa, 0x72, 0x3c, 0xc2, 0x33, 0x26, 0x55,
	0xea, 0x59, 0xb7, 0x3d, 0xf1, 0x54, 0x08, 0xee, 0x8c, 0x0d, 0x56, 0x86, 0xba, 0xb7, 0xb6, 0x4d,
	0x5f, 0x8b, 0xa9, 0xae, 0x65, 0xfb, 0x67, 0xae, 0xc5, 0x6d, 0xfb, 0xb5, 0x02, 0x85, 0xd4, 0x3a,
	0x40, 0x9b, 0x53, 0xc7, 0xd6, 0xbe, 0x27, 0x0d, 0x6d, 0xed, 0x56, 0x9f, 0xbb, 0x4f, 0x6b, 0xe3,
	0xc6, 0x44, 0x8f, 0x6a, 0x3e, 0xda, 0xbb, 0xb8, 0xfa, 0xe8, 0xdf, 0x12, 0x57, 0x28, 0x70, 0x02,
	0xb7, 0x38, 0x2d, 0x4e, 0xc4, 0xd7, 0x63, 0x4d, 0x13, 0xfa, 0x11, 0x62, 0xf3, 0x57, 0x11, 0xcf,
	0xfc, 0x51, 0xc4, 0x73, 0x0f, 0x42, 0x3c, 0xbc, 0x36, 0xd0, 0x1a, 0xe1, 0xc9, 0x7d, 0x21, 0x8f,
	0x8d, 0xf7, 0x83, 0x5c, 0x8e, 0x78, 0x8c, 0x59, 0x64, 0x73, 0x11, 0x39, 0x11, 0x30, 0xf5, 0x93,
	0xe7, 0x68, 0x09, 0x8f, 0x68, 0x5a, 0xfb, 0x6f, 0xf2, 0x52, 0x2f, 0xbf, 0x99, 0x2b, 0x6f, 0x94,
	0xf1, 0x6c, 0x67, 0x6c, 0x3a, 0x1b, 0x64, 0x92, 0x1f, 0xc6, 0x67, 0xa7, 0xda, 0xe4, 0xcf, 0xa9,
	0x5e, 0xcf, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xbf, 0xdc, 0x4c, 0xe2, 0x06, 0x00, 0x00,
}
