// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: response_writer.gen.go.tmpl

package storage

import (
	"github.com/influxdata/influxdb/tsdb"
)

func (w *responseWriter) streamFloatPoints(cur tsdb.FloatBatchCursor) {
	w.sf.DataType = DataTypeFloat
	ss := len(w.res.Frames) - 1

	frame := &ReadResponse_FloatPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]float64, 0, batchSize)}
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_FloatPoints{frame}})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			frame = &ReadResponse_FloatPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]float64, 0, batchSize)}
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_FloatPoints{frame}})
		}
	}
	cur.Close()

	seriesValueCount += b
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		// no points collected, strip series frame
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
		w.vc += seriesValueCount
	}
}

func (w *responseWriter) streamIntegerPoints(cur tsdb.IntegerBatchCursor) {
	w.sf.DataType = DataTypeInteger
	ss := len(w.res.Frames) - 1

	frame := &ReadResponse_IntegerPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]int64, 0, batchSize)}
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_IntegerPoints{frame}})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			frame = &ReadResponse_IntegerPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]int64, 0, batchSize)}
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_IntegerPoints{frame}})
		}
	}
	cur.Close()

	seriesValueCount += b
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		// no points collected, strip series frame
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
		w.vc += seriesValueCount
	}
}

func (w *responseWriter) streamUnsignedPoints(cur tsdb.UnsignedBatchCursor) {
	w.sf.DataType = DataTypeUnsigned
	ss := len(w.res.Frames) - 1

	frame := &ReadResponse_UnsignedPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]uint64, 0, batchSize)}
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_UnsignedPoints{frame}})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			frame = &ReadResponse_UnsignedPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]uint64, 0, batchSize)}
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_UnsignedPoints{frame}})
		}
	}
	cur.Close()

	seriesValueCount += b
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		// no points collected, strip series frame
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
		w.vc += seriesValueCount
	}
}

func (w *responseWriter) streamStringPoints(cur tsdb.StringBatchCursor) {
	w.sf.DataType = DataTypeString
	ss := len(w.res.Frames) - 1

	frame := &ReadResponse_StringPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]string, 0, batchSize)}
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_StringPoints{frame}})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			frame = &ReadResponse_StringPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]string, 0, batchSize)}
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_StringPoints{frame}})
		}
	}
	cur.Close()

	seriesValueCount += b
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		// no points collected, strip series frame
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
		w.vc += seriesValueCount
	}
}

func (w *responseWriter) streamBooleanPoints(cur tsdb.BooleanBatchCursor) {
	w.sf.DataType = DataTypeBoolean
	ss := len(w.res.Frames) - 1

	frame := &ReadResponse_BooleanPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]bool, 0, batchSize)}
	w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_BooleanPoints{frame}})

	var (
		seriesValueCount = 0
		b                = 0
	)

	for {
		ts, vs := cur.Next()
		if len(ts) == 0 {
			break
		}

		frame.Timestamps = append(frame.Timestamps, ts...)
		frame.Values = append(frame.Values, vs...)

		b = len(frame.Timestamps)
		if b >= batchSize {
			seriesValueCount += b
			b = 0
			w.sz += frame.Size()
			if w.sz >= writeSize {
				w.flushFrames()
				if w.err != nil {
					break
				}
			}

			frame = &ReadResponse_BooleanPointsFrame{Timestamps: make([]int64, 0, batchSize), Values: make([]bool, 0, batchSize)}
			w.res.Frames = append(w.res.Frames, ReadResponse_Frame{&ReadResponse_Frame_BooleanPoints{frame}})
		}
	}
	cur.Close()

	seriesValueCount += b
	if seriesValueCount == 0 {
		w.sz -= w.sf.Size()
		// no points collected, strip series frame
		w.res.Frames = w.res.Frames[:ss]
	} else if w.sz > writeSize {
		w.flushFrames()
		w.vc += seriesValueCount
	}
}
